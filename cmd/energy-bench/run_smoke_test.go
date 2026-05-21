//go:build smoke

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRun_SmokeIdleForeground(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("smoke test targets darwin top(1)")
	}
	tmp := t.TempDir()
	o := opts{
		binary:      "/bin/cat",
		scenario:    "idle-foreground",
		fixture:     "small",
		contexts:    1,
		repetitions: 1,
		durationSec: 5,
		outDir:      tmp,
	}
	if err := run(o); err != nil {
		t.Fatalf("run: %v", err)
	}
	entries, err := os.ReadDir(tmp)
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("no report directory created")
	}
	reportPath := filepath.Join(tmp, entries[0].Name(), "report.json")
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	var r Report
	if err := json.Unmarshal(data, &r); err != nil {
		t.Fatalf("unmarshal report: %v", err)
	}
	if r.Primary.WakeupsPerSecond < 0 {
		t.Errorf("nonsensical wakeups_per_s = %v", r.Primary.WakeupsPerSecond)
	}
	// /bin/cat is not an lfk build with the probe wrapper, so probe data
	// will be absent. That's fine for this smoke test — we're verifying
	// the pipeline does not crash and produces a top-derived primary block.
}
