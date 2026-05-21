package energy

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTick_WhenDisabled_IsPlainTeaTick(t *testing.T) {
	t.Setenv("LFK_ENERGY_PROBE", "")
	p, err := Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })

	SetGlobal(p)
	t.Cleanup(func() { SetGlobal(nil) })

	cmd := Tick("test-label", 10*time.Millisecond, func(now time.Time) tea.Msg {
		return now
	})
	require.NotNil(t, cmd)
	msg := cmd()
	_, ok := msg.(time.Time)
	assert.True(t, ok, "wrapper must forward the producer's return value")
}

func TestTick_WhenEnabled_IncrementsCounter(t *testing.T) {
	p, err := startForTest(time.Second, defaultRingCapacity)
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })

	SetGlobal(p)
	t.Cleanup(func() { SetGlobal(nil) })

	cmd := Tick("alpha", 5*time.Millisecond, func(now time.Time) tea.Msg { return now })
	_ = cmd()

	counts := p.snapshotTickCounts()
	assert.EqualValues(t, 1, counts["alpha"])
}

func TestTick_NilGlobal_IsSafe(t *testing.T) {
	SetGlobal(nil)
	cmd := Tick("x", time.Millisecond, func(now time.Time) tea.Msg { return now })
	require.NotNil(t, cmd)
	msg := cmd()
	_, ok := msg.(time.Time)
	assert.True(t, ok)
}
