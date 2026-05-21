package energy

import (
	"testing"
	"time"

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
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	assert.True(t, p.Enabled())
}

func TestProbe_StopIsIdempotent(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "1")
	p, err := Start()
	require.NoError(t, err)
	assert.NoError(t, p.Stop())
	assert.NoError(t, p.Stop(), "Stop must be safe to call twice")
}

func TestProbe_DisabledStartHasNoGoroutine(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "")
	before := runtimeNumGoroutine()
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	// Give a hypothetical leaked goroutine a moment to register before
	// snapshotting; if Start spawned anything it will be visible here.
	time.Sleep(20 * time.Millisecond)
	after := runtimeNumGoroutine()
	assert.Equal(t, before, after,
		"disabled probe must not start a goroutine")
}

func TestProbe_RunIDIsStable(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "1")
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })
	id := p.RunID()
	assert.NotEmpty(t, id)
	assert.Equal(t, id, p.RunID(), "RunID must be stable across calls")
}

func runtimeNumGoroutine() int { return numGoroutineForTest() }
