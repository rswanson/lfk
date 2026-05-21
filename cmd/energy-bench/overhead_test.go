package main

import (
	"context"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestProbe_OverheadUnder1Percent runs a stand-in process twice — once
// with the probe enabled, once without — and asserts that the probe adds
// less than 1% to wakeups_per_s.
//
// This is slow (two ~15s runs) and macOS-only. Skipped in -short mode.
func TestProbe_OverheadUnder1Percent(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS-only")
	}
	if testing.Short() {
		t.Skip("slow test; skipped in short mode")
	}

	bin := buildOverheadStandin(t)
	measure := func(env []string) float64 {
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		c := exec.CommandContext(ctx, bin)
		c.Env = env
		require.NoError(t, c.Start())
		defer func() { _ = c.Process.Kill(); _ = c.Wait() }()
		time.Sleep(2 * time.Second) // warmup

		out, _ := exec.Command("top", "-l", "5", "-stats", "pid,cpu,power,idlew", "-pid", strconv.Itoa(c.Process.Pid)).Output()
		samples, err := parseTopOutputForPID(string(out), c.Process.Pid)
		require.NoError(t, err)
		var avg float64
		for _, x := range samples {
			avg += x.WakeupsPerS
		}
		if len(samples) > 0 {
			avg /= float64(len(samples))
		}
		return avg
	}

	off := measure([]string{"LFK_ENERGY_PROBE=", "PATH=" + os.Getenv("PATH")})
	on := measure([]string{"LFK_ENERGY_PROBE=1", "LFK_DATA_DIR=" + t.TempDir(), "PATH=" + os.Getenv("PATH")})

	t.Logf("wakeups_per_s: off=%.2f on=%.2f", off, on)
	if off == 0 {
		t.Skip("baseline wakeups_per_s is 0; cannot compute overhead ratio")
	}
	overhead := (on - off) / off
	assert.Less(t, overhead, 0.01, "probe must add <1%% to wakeups_per_s (got %.4f)", overhead)
}

func buildOverheadStandin(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	out := dir + "/standin"
	c := exec.Command("go", "build", "-o", out, "github.com/janosmiko/lfk/internal/perf/energy/probe_overhead_main")
	c.Stdout = nil
	c.Stderr = nil
	require.NoError(t, c.Run())
	return out
}
