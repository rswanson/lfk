package main

import (
	"context"
	"time"
)

type scriptedStep struct {
	key     string
	repeat  int
	pauseMs int
}

type expanded struct {
	out     string
	pauseMs int
}

var keyMap = map[string]string{
	"enter": "\r",
	"tab":   "\t",
	"esc":   "\x1b",
	"down":  "\x1b[B",
	"up":    "\x1b[A",
	"right": "\x1b[C",
	"left":  "\x1b[D",
}

func keyToBytes(k string) string {
	if v, ok := keyMap[k]; ok {
		return v
	}
	return k
}

func expandScript(steps []scriptedStep) []expanded {
	out := make([]expanded, 0, len(steps))
	for _, s := range steps {
		n := s.repeat
		if n <= 0 {
			n = 1
		}
		for i := 0; i < n; i++ {
			out = append(out, expanded{out: keyToBytes(s.key), pauseMs: s.pauseMs})
		}
	}
	return out
}

// defaultActiveScript: open pods -> down 20 -> enter (open logs) -> esc -> tab (next context) -> repeat.
func defaultActiveScript() []scriptedStep {
	return []scriptedStep{
		{key: "down", repeat: 20, pauseMs: 50},
		{key: "enter", pauseMs: 500},
		{key: "esc", pauseMs: 200},
		{key: "tab", pauseMs: 200},
	}
}

func runActiveScripted(parent context.Context, cfg scenarioConfig) error {
	ctx, cancel := context.WithTimeout(parent, time.Duration(cfg.durationSec+5)*time.Second)
	defer cancel()

	sess, err := startPTY(ctx, cfg.binary, cfg.args, cfg.env)
	if err != nil {
		return err
	}
	defer sess.Close()

	time.Sleep(500 * time.Millisecond) // let UI render
	script := expandScript(defaultActiveScript())
	end := time.Now().Add(time.Duration(cfg.durationSec) * time.Second)
	i := 0
	for time.Now().Before(end) {
		step := script[i%len(script)]
		if err := sess.SendString(step.out); err != nil {
			return err
		}
		time.Sleep(time.Duration(step.pauseMs) * time.Millisecond)
		i++
	}
	return nil
}
