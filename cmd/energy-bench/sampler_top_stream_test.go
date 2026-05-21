package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

// fakeTopScript writes a shell script to dir that emits a deterministic
// top-like stream: two PID rows per block, 1s sleep between blocks,
// PIDs 1234 and 5678. IDLEW columns increment each block.
func fakeTopScript(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "fake-top.sh")
	script := `#!/bin/sh
i=0
while :; do
  echo "Processes: ..."
  echo "PID    %CPU   POWER  IDLEW"
  echo "1234   1.0    0.5    $((10 + i))"
  echo "5678   2.0    1.0    $((20 + i))"
  i=$((i + 1))
  sleep 1
done
`
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatalf("write fake-top: %v", err)
	}
	return path
}

func TestTopStreamSampler_FiltersToPID(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("sampler is darwin/linux only")
	}
	dir := t.TempDir()
	cmd := fakeTopScript(t, dir)

	s := newTopStreamSampler(cmd, []string{}, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Start(ctx, 1234); err != nil {
		t.Fatalf("Start: %v", err)
	}
	// Let it collect at least 3 sample blocks (warmup + 2 kept).
	time.Sleep(3500 * time.Millisecond)
	got := s.Stop()
	if len(got) < 2 {
		t.Fatalf("want >=2 samples for PID 1234, got %d", len(got))
	}
	for _, sm := range got {
		if sm.PID != 1234 {
			t.Errorf("unexpected PID in filtered sample: %d", sm.PID)
		}
	}
}

func TestTopStreamSampler_StopIsIdempotent(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("sampler is darwin/linux only")
	}
	dir := t.TempDir()
	cmd := fakeTopScript(t, dir)
	s := newTopStreamSampler(cmd, []string{}, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Start(ctx, 1234); err != nil {
		t.Fatalf("Start: %v", err)
	}
	_ = s.Stop()
	_ = s.Stop() // must not panic or block
}
