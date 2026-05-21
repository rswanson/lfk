// Package energy provides an opt-in, low-overhead in-process probe for
// measuring lfk's runtime behaviour during energy benchmarks. It is
// activated by LFK_ENERGY_PROBE=1 and is otherwise a no-op (no goroutines,
// no allocations on hot paths).
package energy

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"runtime"
	"sync"
)

const envEnable = "LFK_ENERGY_PROBE"

// Probe is the running probe handle. Stop is idempotent.
type Probe struct {
	enabled bool
	runID   string

	mu      sync.Mutex
	stopped bool
	stopCh  chan struct{}
}

// Start returns an initialised probe. When LFK_ENERGY_PROBE is unset or
// empty, the returned probe is disabled and Stop is a no-op.
func Start() (*Probe, error) {
	enabled := os.Getenv(envEnable) == "1"
	p := &Probe{enabled: enabled}
	if !enabled {
		return p, nil
	}
	id, err := newRunID()
	if err != nil {
		return nil, err
	}
	p.runID = id
	p.stopCh = make(chan struct{})
	// Sampling and output goroutines are added in later tasks.
	return p, nil
}

// Enabled reports whether the probe is active.
func (p *Probe) Enabled() bool { return p != nil && p.enabled }

// RunID returns the unique identifier for this probe run.
func (p *Probe) RunID() string { return p.runID }

// Stop shuts down the probe. It is safe to call more than once.
func (p *Probe) Stop() error {
	if p == nil || !p.enabled {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.stopped {
		return nil
	}
	p.stopped = true
	close(p.stopCh)
	return nil
}

func newRunID() (string, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", errors.New("energy: cannot generate run id")
	}
	return hex.EncodeToString(b[:]), nil
}

// numGoroutineForTest is a thin wrapper used by the test file to avoid
// importing "runtime" there.
func numGoroutineForTest() int { return runtime.NumGoroutine() }
