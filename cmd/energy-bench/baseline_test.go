package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBaseline_WriteAndRead(t *testing.T) {
	dir := t.TempDir()
	r := Report{Scenario: "idle-foreground", Primary: primaryMetrics{WakeupsPerSecond: 42}}
	require.NoError(t, writeBaseline(filepath.Join(dir, "idle-foreground.json"), r))

	got, err := readBaseline(filepath.Join(dir, "idle-foreground.json"))
	require.NoError(t, err)
	assert.Equal(t, "idle-foreground", got.Scenario)
	assert.Equal(t, 42.0, got.Primary.WakeupsPerSecond)
}

func TestBaseline_MissingFileReturnsNil(t *testing.T) {
	got, err := readBaseline("/does/not/exist.json")
	require.NoError(t, err)
	assert.Nil(t, got)
}
