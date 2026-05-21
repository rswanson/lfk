package energy

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlush_WritesJSONLToDataDir(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("LFK_DATA_DIR", dir)

	p, err := startForTest(10*time.Millisecond, defaultRingCapacity)
	require.NoError(t, err)
	time.Sleep(60 * time.Millisecond)
	require.NoError(t, p.Stop())

	// paths.DataDir() returns LFK_DATA_DIR verbatim, so our code appends "lfk"
	// and "energy" — resulting in dir/lfk/energy/*.jsonl.
	matches, err := filepath.Glob(filepath.Join(dir, "lfk", "energy", "*.jsonl"))
	require.NoError(t, err)
	if len(matches) == 0 {
		// Alternative layout: paths.DataDir() already returned dir/lfk and our
		// code joined only "energy".
		matches, err = filepath.Glob(filepath.Join(dir, "energy", "*.jsonl"))
		require.NoError(t, err)
	}
	require.Len(t, matches, 1, "expected exactly one JSONL file under the data dir")

	f, err := os.Open(matches[0])
	require.NoError(t, err)
	defer f.Close()

	count := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var s Sample
		require.NoError(t, json.Unmarshal(sc.Bytes(), &s))
		assert.Positive(t, s.WallNanos)
		count++
	}
	assert.GreaterOrEqual(t, count, 3)
}

func TestFlush_DisabledProbeIsNoOp(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("LFK_ENERGY_PROBE", "")
	t.Setenv("LFK_DATA_DIR", dir)

	p, err := Start()
	require.NoError(t, err)
	require.NoError(t, p.Stop())

	matches, _ := filepath.Glob(filepath.Join(dir, "lfk", "energy", "*.jsonl"))
	more, _ := filepath.Glob(filepath.Join(dir, "energy", "*.jsonl"))
	matches = append(matches, more...)
	assert.Empty(t, matches, "disabled probe must not create files")
}
