package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestE2E_IdleForegroundProducesReport runs the full pipeline against a
// trivial stand-in process (sh -c sleep 2) so we can verify the harness
// produces a report.json + report.md without needing a built lfk binary.
func TestE2E_IdleForegroundProducesReport(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("scenario depends on macOS top(1); skipping on non-darwin")
	}
	if _, err := exec.LookPath("top"); err != nil {
		t.Skip("no top available")
	}
	outDir := t.TempDir()
	o := opts{
		binary:      "/bin/sh",
		scenario:    "idle-foreground",
		fixture:     "small",
		contexts:    3,
		repetitions: 1,
		durationSec: 2,
		outDir:      outDir,
	}
	require.NoError(t, run(o))

	matches, err := filepath.Glob(filepath.Join(outDir, "*", "report.md"))
	require.NoError(t, err)
	require.Len(t, matches, 1)
	data, err := os.ReadFile(matches[0])
	require.NoError(t, err)
	assert.Contains(t, string(data), "idle-foreground")
}
