package energy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/janosmiko/lfk/internal/paths"
)

// flushMu serializes flush() to make concurrent SIGUSR1 + Stop safe.
// Without this, two flushes can race on os.Create() for the same path.
var flushMu sync.Mutex

// flush writes the entire ring buffer to a JSONL file under the data dir.
// Called by Stop and (via FlushOnSignal) on SIGUSR1. Safe to call concurrently.
func (p *Probe) flush() error {
	if p == nil || !p.enabled {
		return nil
	}
	flushMu.Lock()
	defer flushMu.Unlock()

	dir, err := paths.DataDir()
	if err != nil {
		return fmt.Errorf("energy: data dir: %w", err)
	}
	// paths.DataDir() always returns the lfk data directory: it appends
	// "lfk" to XDG_DATA_HOME / the OS default, and treats LFK_DATA_DIR as
	// the literal lfk directory. So we only add the "energy" subdir here.
	outDir := filepath.Join(dir, "energy")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("energy: mkdir: %w", err)
	}
	path := filepath.Join(outDir, p.runID+".jsonl")
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("energy: create: %w", err)
	}
	defer func() { _ = f.Close() }()
	enc := json.NewEncoder(f)
	for _, s := range p.snapshotSamples() {
		if err := enc.Encode(&s); err != nil {
			return fmt.Errorf("energy: encode: %w", err)
		}
	}
	return nil
}

// FlushOnSignal writes the current ring buffer without stopping the sampler.
// Safe to call repeatedly (e.g., on SIGUSR1) and concurrent with Stop.
func (p *Probe) FlushOnSignal() error {
	if p == nil || !p.enabled {
		return nil
	}
	return p.flush()
}
