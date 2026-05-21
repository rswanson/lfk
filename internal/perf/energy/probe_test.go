package energy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProbe_DisabledByDefault(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "")
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	assert.False(t, p.Enabled(), "probe must be disabled when env var is unset")
}

func TestProbe_EnabledByEnvVar(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "1")
	t.Setenv("LFK_DATA_DIR", t.TempDir())
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	assert.True(t, p.Enabled())
}

func TestProbe_StopIsIdempotent(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "1")
	t.Setenv("LFK_DATA_DIR", t.TempDir())
	p, err := Start()
	require.NoError(t, err)
	assert.NoError(t, p.Stop())
	assert.NoError(t, p.Stop(), "Stop must be safe to call twice")
}

func TestProbe_DisabledStartHasNoGoroutine(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "")
	before := readGoroutineCount(t)
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	assert.LessOrEqual(t, readGoroutineCount(t), before,
		"disabled probe must not start a goroutine")
}

func readGoroutineCount(t *testing.T) int {
	t.Helper()
	return runtimeNumGoroutine()
}

func TestProbe_RunIDIsStable(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "1")
	t.Setenv("LFK_DATA_DIR", t.TempDir())
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	id := p.RunID()
	assert.NotEmpty(t, id)
	assert.Equal(t, id, p.RunID(), "RunID must be stable across calls")
}

func runtimeNumGoroutine() int { return numGoroutineForTest() }

var _ = os.Getenv // keep "os" used in this file
