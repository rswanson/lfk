package energy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSampler_CapturesAtLeastOneSampleIn200ms(t *testing.T) {
	p, err := startForTest(20*time.Millisecond, defaultRingCapacity)
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })

	time.Sleep(200 * time.Millisecond)

	samples := p.snapshotSamples()
	assert.GreaterOrEqual(t, len(samples), 5, "expected at least 5 samples in 200ms at 20ms cadence")
	s := samples[0]
	assert.Positive(t, s.Goroutines)
	assert.NotZero(t, s.WallNanos)
}

func TestSampler_RingBufferDropsOldest(t *testing.T) {
	p, err := startForTest(10*time.Millisecond, 4)
	require.NoError(t, err)
	t.Cleanup(func() { _ = p.Stop() })

	time.Sleep(150 * time.Millisecond)

	samples := p.snapshotSamples()
	assert.LessOrEqual(t, len(samples), 4, "ring must cap at configured capacity")
	for i := 1; i < len(samples); i++ {
		assert.GreaterOrEqual(t, samples[i].WallNanos, samples[i-1].WallNanos)
	}
}
