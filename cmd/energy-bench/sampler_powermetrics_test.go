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

func TestPowermetrics_M4FormatHWOnlyMultiPCluster(t *testing.T) {
	// M-series (M4 seen in the wild) emits ONLY "<X>{N}-Cluster HW active
	// residency:" lines and may have multiple P-clusters (P0, P1). Package
	// Power is absent. Verify the parser averages multi-P-cluster residency
	// and tolerates a missing Package Power line.
	data, err := os.ReadFile("testdata/powermetrics_sample_m4.txt")
	require.NoError(t, err)
	pm, err := parsePowermetricsOutput(string(data))
	require.NoError(t, err)
	// E-Cluster HW: 50.00%
	assert.InDelta(t, 50.0, pm.ECoreActiveResidency, 0.01)
	// (P0 HW 10.00 + P1 HW 20.00) / 2 = 15.00
	assert.InDelta(t, 15.0, pm.PCoreActiveResidency, 0.01)
	// No "Package Power:" line in the fixture.
	assert.Equal(t, 0.0, pm.PackagePowerMW)
}
