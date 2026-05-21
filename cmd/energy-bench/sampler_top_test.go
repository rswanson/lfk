package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTopSampler_ParsesRecordedSample(t *testing.T) {
	data, err := os.ReadFile("testdata/top_sample.txt")
	require.NoError(t, err)
	samples, err := parseTopOutput(string(data))
	require.NoError(t, err)
	require.NotEmpty(t, samples)
	s := samples[0]
	assert.GreaterOrEqual(t, s.CPUPercent, 0.0)
	assert.GreaterOrEqual(t, s.WakeupsPerS, 0.0)
}

func TestTopSampler_FiltersToPID(t *testing.T) {
	data, err := os.ReadFile("testdata/top_sample.txt")
	require.NoError(t, err)
	// Find any PID in the data and assert filtering returns only rows for it.
	all, err := parseTopOutput(string(data))
	require.NoError(t, err)
	require.NotEmpty(t, all)
	pid := all[0].PID
	samples, err := parseTopOutputForPID(string(data), pid)
	require.NoError(t, err)
	for _, s := range samples {
		assert.Equal(t, pid, s.PID)
	}
}

func TestTopSampler_EmptyInputError(t *testing.T) {
	_, err := parseTopOutput("")
	assert.Error(t, err)
}

func TestTopSampler_SkipsUnparsableRows(t *testing.T) {
	// A header block with one valid row and one unparsable row.
	input := "Processes: 1 total\n\nPID    %CPU CSW\n12345  1.5  100\nbadrow\n"
	samples, err := parseTopOutput(input)
	require.NoError(t, err)
	require.Len(t, samples, 1)
	assert.Equal(t, 12345, samples[0].PID)
	assert.InDelta(t, 1.5, samples[0].CPUPercent, 0.001)
	assert.InDelta(t, 100.0, samples[0].WakeupsPerS, 0.001)
}

func TestTopSampler_HandlesTrailingPlusInCSW(t *testing.T) {
	// CSW values suffixed with '+' (cumulative overflow marker) should parse cleanly.
	input := "Processes: 1 total\n\nPID    %CPU CSW\n617    32.6 710558144+\n"
	samples, err := parseTopOutput(input)
	require.NoError(t, err)
	require.Len(t, samples, 1)
	assert.InDelta(t, 710558144.0, samples[0].WakeupsPerS, 0.001)
}

func TestTopSampler_FiltersToPIDEmpty(t *testing.T) {
	input := "Processes: 1 total\n\nPID    %CPU CSW\n12345  1.5  100\n"
	samples, err := parseTopOutputForPID(input, 99999)
	require.NoError(t, err)
	assert.Empty(t, samples)
}
