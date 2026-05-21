package main

import (
	"context"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

// TestScenarios_FireOnProcessStart verifies that every scenario invokes
// cfg.onProcessStart exactly once with a positive PID after the subprocess
// starts. Uses /bin/cat as a stand-in for lfk: it stays alive long enough
// to be probed and exits cleanly when the PTY closes.
func TestScenarios_FireOnProcessStart(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("PTY scenarios are darwin/linux only")
	}
	t.Setenv("LFK_TEST_DISABLE_UNFOCUS", "1")
	// On Linux, runIdleBackground's unfocusTerminal is already a no-op
	// (non-darwin guard); on macOS the LFK_TEST_DISABLE_UNFOCUS env var
	// set above suppresses the osascript call.
	cases := []struct {
		name string
		run  func(context.Context, scenarioConfig) error
	}{
		{"idle-foreground", runIdleForeground},
		{"idle-background", func(ctx context.Context, c scenarioConfig) error {
			return runIdleBackground(ctx, c, "Ghostty")
		}},
		{"active-scripted", runActiveScripted},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			var gotPID atomic.Int64
			cfg := scenarioConfig{
				binary:      "/bin/cat",
				args:        nil,
				env:         []string{},
				durationSec: 1,
				onProcessStart: func(pid int) {
					calls.Add(1)
					gotPID.Store(int64(pid))
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			t.Cleanup(cancel)
			if err := tc.run(ctx, cfg); err != nil {
				t.Fatalf("scenario error: %v", err)
			}
			if got := calls.Load(); got != 1 {
				t.Fatalf("onProcessStart called %d times, want 1", got)
			}
			if got := gotPID.Load(); got <= 0 {
				t.Fatalf("onProcessStart got pid=%d, want >0", got)
			}
		})
	}
}
