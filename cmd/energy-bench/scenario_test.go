package main

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPTY_RunsAndCapturesOutput(t *testing.T) {
	if _, err := exec.LookPath("cat"); err != nil {
		t.Skip("no cat binary available")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	sess, err := startPTY(ctx, "cat", nil, nil)
	require.NoError(t, err)
	defer sess.Close()

	require.NoError(t, sess.SendString("hello\n"))
	out, err := sess.ReadUntil("hello", time.Second)
	require.NoError(t, err)
	assert.Contains(t, out, "hello")
}

func TestScenarioIdleForeground_LifecycleNoCrash(t *testing.T) {
	cfg := scenarioConfig{
		binary:      "/bin/sh",
		args:        []string{"-c", "sleep 1"},
		env:         nil,
		durationSec: 1,
	}
	err := runIdleForeground(context.Background(), cfg)
	assert.NoError(t, err)
}

func TestScenarioIdleBackground_BuildsOsascriptArgs(t *testing.T) {
	args := osascriptUnfocusArgs("Ghostty")
	require.Len(t, args, 2)
	assert.Equal(t, "-e", args[0])
	assert.Contains(t, args[1], "Finder")
}

func TestScriptedSteps_AreInterpreted(t *testing.T) {
	steps := []scriptedStep{
		{key: "j", repeat: 3, pauseMs: 50},
		{key: "enter", pauseMs: 100},
	}
	exp := expandScript(steps)
	assert.Len(t, exp, 4)
	assert.Equal(t, "j", exp[0].out)
	assert.Equal(t, "\r", exp[3].out)
}
