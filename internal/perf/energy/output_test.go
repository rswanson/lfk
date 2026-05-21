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

	// paths.DataDir() treats LFK_DATA_DIR as the lfk data dir verbatim;
	// the flush appends only "energy". The output path is dir/energy/*.jsonl.
	matches, err := filepath.Glob(filepath.Join(dir, "energy", "*.jsonl"))
	require.NoError(t, err)
	require.Len(t, matches, 1, "expected exactly one JSONL file at dir/energy/")
	// Defensive: confirm we did NOT write to the historical buggy dir/lfk/energy.
	wrong, _ := filepath.Glob(filepath.Join(dir, "lfk", "energy", "*.jsonl"))
	assert.Empty(t, wrong, "must not double-append lfk segment")

	f, err := os.Open(matches[0])
	require.NoError(t, err)
	defer func() { _ = f.Close() }()

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

	matches, _ := filepath.Glob(filepath.Join(dir, "energy", "*.jsonl"))
	assert.Empty(t, matches, "disabled probe must not create files")
}
