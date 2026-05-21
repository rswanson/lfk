package energy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/janosmiko/lfk/internal/paths"
)

// flush writes the entire ring buffer to a JSONL file under the data dir.
// Called by Stop and (via FlushOnSignal) on SIGUSR1.
func (p *Probe) flush() error {
	if p == nil || !p.enabled {
		return nil
	}
	dir, err := paths.DataDir()
	if err != nil {
		return fmt.Errorf("energy: data dir: %w", err)
	}
	// paths.DataDir() returns LFK_DATA_DIR verbatim (no "lfk" appended) when
	// the env var is set. When XDG or the OS default is used it already
	// appends "lfk". We always add "lfk" here so that the verbatim-env-var
	// case is handled correctly; when the XDG/default path is used the caller
	// has not set LFK_DATA_DIR so the variable does not affect resolution.
	// NOTE: this intentionally produces .../lfk/energy/... in all cases.
	outDir := filepath.Join(dir, "lfk", "energy")
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
// Safe to call repeatedly (e.g., on SIGUSR1).
func (p *Probe) FlushOnSignal() error {
	if p == nil || !p.enabled {
		return nil
	}
	return p.flush()
}
