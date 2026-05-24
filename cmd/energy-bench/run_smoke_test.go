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
	// /bin/cat is not an lfk build with the probe wrapper, so probe
	// data is absent and the PID filter matches no rows (cat is not a
	// long-lived foreground process at the system top level). What we
	// can assert is that the pipeline completed and produced a
	// structurally valid report at the expected scenario.
	if r.Scenario != "idle-foreground" {
		t.Errorf("got scenario %q, want idle-foreground", r.Scenario)
	}
	if r.Reps != 1 {
		t.Errorf("got reps=%d, want 1 (smoke test passes -reps 1 implicitly via default)", r.Reps)
	}
}
