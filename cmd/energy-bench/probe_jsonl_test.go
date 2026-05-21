package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/janosmiko/lfk/internal/perf/energy"
)

func writeJSONLFixture(t *testing.T, dir string, samples []energy.Sample) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	path := filepath.Join(dir, "run-abc.jsonl")
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = f.Close() }()
	enc := json.NewEncoder(f)
	for _, s := range samples {
		if err := enc.Encode(&s); err != nil {
			t.Fatalf("encode: %v", err)
		}
	}
}

func TestIngestProbe_AggregatesTickRatesAndDiagnostics(t *testing.T) {
	dir := t.TempDir()
	energyDir := filepath.Join(dir, "energy")

	// 11 samples, 1s apart. Tick counts climb linearly so per-second
	// rates are deterministic. The first sample's tick counts are the
	// "zero" (treated as a baseline by the aggregator).
	samples := make([]energy.Sample, 11)
	for i := range 11 {
		samples[i] = energy.Sample{
			WallNanos:       int64(i) * int64(1e9),
			Goroutines:      10 + i,
			HeapAllocBytes:  uint64(10_000_000 + i*1_000_000),
			NumGC:           uint32(i),
			GCPauseTotalNs:  uint64(i) * 1_000_000,
			GCCPUSeconds:    float64(i) * 0.01, // +0.1s of GC CPU over the 10s window
			SchedLatencyP99: 0.0005,            // 500us
			TickCounts: map[string]uint64{
				"pods-refresh":          uint64(i * 2),  // 2/s
				"kubetris-render-100ms": uint64(i * 10), // 10/s
			},
		}
	}
	writeJSONLFixture(t, energyDir, samples)

	agg, err := ingestProbe(energyDir)
	if err != nil {
		t.Fatalf("ingestProbe: %v", err)
	}
	if agg == nil {
		t.Fatal("ingestProbe returned nil")
	}
	if agg.GoroutinesMax != 20 {
		t.Errorf("GoroutinesMax = %d, want 20", agg.GoroutinesMax)
	}
	// 10s wall, 0.1s of GC CPU => fraction = 0.01
	if got := agg.GCCPUFraction; got < 0.0095 || got > 0.0105 {
		t.Errorf("GCCPUFraction = %v, want ~0.01", got)
	}
	if got := agg.TickRatesPerSecond["pods-refresh"]; got < 1.9 || got > 2.1 {
		t.Errorf("pods-refresh rate = %v, want ~2", got)
	}
	if got := agg.TickRatesPerSecond["kubetris-render-100ms"]; got < 9.9 || got > 10.1 {
		t.Errorf("kubetris-render-100ms rate = %v, want ~10", got)
	}
	if agg.TickCountsTotal["pods-refresh"] != 20 {
		t.Errorf("pods-refresh total = %d, want 20", agg.TickCountsTotal["pods-refresh"])
	}
}

func TestIngestProbe_MissingDirReturnsNil(t *testing.T) {
	dir := t.TempDir() // no energy/ subdir created
	agg, err := ingestProbe(filepath.Join(dir, "energy"))
	if err != nil {
		t.Fatalf("ingestProbe missing dir: %v", err)
	}
	if agg != nil {
		t.Errorf("want nil aggregates for missing dir, got %+v", agg)
	}
}

func TestIngestProbe_SingleSampleReturnsZeroRates(t *testing.T) {
	dir := t.TempDir()
	energyDir := filepath.Join(dir, "energy")
	writeJSONLFixture(t, energyDir, []energy.Sample{{
		WallNanos:    1_000_000_000,
		Goroutines:   5,
		GCCPUSeconds: 0.5,
		TickCounts:   map[string]uint64{"pods-refresh": 100},
	}})
	agg, err := ingestProbe(energyDir)
	if err != nil {
		t.Fatalf("ingestProbe: %v", err)
	}
	if agg == nil {
		t.Fatal("want non-nil aggregates for single sample")
	}
	if len(agg.TickRatesPerSecond) != 0 {
		t.Errorf("want no rates with single sample, got %v", agg.TickRatesPerSecond)
	}
}
