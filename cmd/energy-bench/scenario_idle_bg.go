package main

import (
	"context"
	"os/exec"
	"runtime"
	"time"
)

// osascriptUnfocusArgs returns the args for osascript to send focus away
// from terminalApp. We activate Finder (always available) so the focus
// transfer is robust across Mission Control spaces.
func osascriptUnfocusArgs(terminalApp string) []string {
	_ = terminalApp // currently unused; kept for future per-app routing
	return []string{"-e", "tell application \"Finder\" to activate"}
}

func unfocusTerminal(parent context.Context, terminalApp string) error {
	if runtime.GOOS != "darwin" {
		return nil // best-effort on non-macOS
	}
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	return exec.CommandContext(ctx, "osascript", osascriptUnfocusArgs(terminalApp)...).Run()
}

// runIdleBackground launches the target, blurs the terminal, then waits.
func runIdleBackground(parent context.Context, cfg scenarioConfig, terminalApp string) error {
	ctx, cancel := context.WithTimeout(parent, time.Duration(cfg.durationSec+5)*time.Second)
	defer cancel()

	sess, err := startPTY(ctx, cfg.binary, cfg.args, cfg.env)
	if err != nil {
		return err
	}
	defer sess.Close()

	if cfg.onProcessStart != nil {
		cfg.onProcessStart(sess.PID())
	}

	// Give the target ~1s to start drawing before we unfocus.
	time.Sleep(time.Second)
	if err := unfocusTerminal(ctx, terminalApp); err != nil {
		return err
	}

	select {
	case <-time.After(time.Duration(cfg.durationSec) * time.Second):
	case <-ctx.Done():
	}
	return nil
}
