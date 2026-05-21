package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPowermetrics_ParsesPCoreECoreAndPackagePower(t *testing.T) {
	data, err := os.ReadFile("testdata/powermetrics_sample.txt")
	require.NoError(t, err)
	pm, err := parsePowermetricsOutput(string(data))
	require.NoError(t, err)
	assert.InDelta(t, 12.34, pm.PCoreActiveResidency, 0.01)
	assert.InDelta(t, 24.10, pm.ECoreActiveResidency, 0.01)
	// Package Power appears twice in the testdata; parser should average.
	assert.InDelta(t, 399.5, pm.PackagePowerMW, 0.5)
}

func TestPowermetrics_EmptyInputDoesNotPanic(t *testing.T) {
	pm, err := parsePowermetricsOutput("")
	require.NoError(t, err)
	assert.Equal(t, 0.0, pm.PCoreActiveResidency)
	assert.Equal(t, 0.0, pm.ECoreActiveResidency)
	assert.Equal(t, 0.0, pm.PackagePowerMW)
}
