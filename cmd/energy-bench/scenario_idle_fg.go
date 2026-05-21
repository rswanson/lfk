package main

import (
	"context"
	"time"
)

type scenarioConfig struct {
	binary      string
	args        []string
	env         []string
	durationSec int
}

// runIdleForeground launches the target, waits durationSec, and returns.
// The probe (running inside the target via LFK_ENERGY_PROBE=1) does the
// real measurement; this function only owns the lifecycle.
func runIdleForeground(parent context.Context, cfg scenarioConfig) error {
	ctx, cancel := context.WithTimeout(parent, time.Duration(cfg.durationSec+5)*time.Second)
	defer cancel()

	sess, err := startPTY(ctx, cfg.binary, cfg.args, cfg.env)
	if err != nil {
		return err
	}
	defer sess.Close()

	select {
	case <-time.After(time.Duration(cfg.durationSec) * time.Second):
	case <-ctx.Done():
	}
	return nil
}
