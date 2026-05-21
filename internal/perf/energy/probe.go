// Package energy provides an opt-in, low-overhead in-process probe for
// measuring lfk's runtime behaviour during energy benchmarks. It is
// activated by LFK_ENERGY_PROBE=1 and is otherwise a no-op (no goroutines,
// no allocations on hot paths).
package energy

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"runtime/metrics"
	"sync"
	"time"
)

const envEnable = "LFK_ENERGY_PROBE"

// Probe is the running probe handle. Stop is idempotent.
type Probe struct {
	enabled bool
	runID   string

	mu      sync.Mutex
	stopped bool
	stopCh  chan struct{}
	wg      sync.WaitGroup

	tickInterval time.Duration

	// Per-probe metrics.Sample slices — avoids races when multiple Probe
	// instances call metrics.Read concurrently.
	schedSamples []metrics.Sample
	gcCPUSamples []metrics.Sample

	ringMu  sync.Mutex
	ring    []Sample
	ringCap int

	tickMu     sync.Mutex
	tickCounts map[string]uint64
}

// Start returns an initialised probe. When LFK_ENERGY_PROBE is unset or
// empty, the returned probe is disabled and Stop is a no-op.
func Start() (*Probe, error) {
	enabled := os.Getenv(envEnable) == "1"
	p := &Probe{
		enabled:      enabled,
		tickInterval: time.Second,
		ringCap:      defaultRingCapacity,
		tickCounts:   make(map[string]uint64),
	}
	if !enabled {
		return p, nil
	}
	id, err := newRunID()
	if err != nil {
		return nil, err
	}
	p.runID = id
	p.stopCh = make(chan struct{})
	p.schedSamples = []metrics.Sample{{Name: "/sched/latency:seconds"}}
	p.gcCPUSamples = []metrics.Sample{{Name: "/cpu/classes/gc/total:cpu-seconds"}}
	p.startSampler()
	return p, nil
}

// startForTest builds an enabled probe with custom tick interval and ring
// capacity for tests. Equivalent to Start() but with injectable config so
// the sampler goroutine reads the test values before starting.
func startForTest(tickInterval time.Duration, ringCap int) (*Probe, error) {
	id, err := newRunID()
	if err != nil {
		return nil, err
	}
	p := &Probe{
		enabled:      true,
		runID:        id,
		tickInterval: tickInterval,
		ringCap:      ringCap,
		tickCounts:   make(map[string]uint64),
		stopCh:       make(chan struct{}),
	}
	p.schedSamples = []metrics.Sample{{Name: "/sched/latency:seconds"}}
	p.gcCPUSamples = []metrics.Sample{{Name: "/cpu/classes/gc/total:cpu-seconds"}}
	p.startSampler()
	return p, nil
}

// Enabled reports whether the probe is active.
func (p *Probe) Enabled() bool { return p != nil && p.enabled }

// RunID returns the unique identifier for this probe run.
func (p *Probe) RunID() string { return p.runID }

// Stop shuts down the probe and waits for the sampler goroutine to exit.
// It is safe to call more than once.
func (p *Probe) Stop() error {
	if p == nil || !p.enabled {
		return nil
	}
	p.mu.Lock()
	if p.stopped {
		p.mu.Unlock()
		return nil
	}
	p.stopped = true
	close(p.stopCh)
	p.mu.Unlock()
	p.wg.Wait()
	return p.flush()
}

func newRunID() (string, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("energy: cannot generate run id: %w", err)
	}
	return hex.EncodeToString(b[:]), nil
}

// numGoroutineForTest is a thin wrapper used by the test file to avoid
// importing "runtime" there.
func numGoroutineForTest() int { return runtime.NumGoroutine() }
