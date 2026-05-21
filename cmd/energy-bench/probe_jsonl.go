package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/janosmiko/lfk/internal/perf/energy"
)

const (
	nanosPerSecond = 1e9
	bytesPerMB     = 1e6
	secsToMicros   = 1e6
)

// ProbeAggregates is the lfk-specific portion of the report, derived from
// the in-process probe's JSONL output.
type ProbeAggregates struct {
	WallSeconds        float64            `json:"wall_seconds"`
	GoroutinesMean     float64            `json:"goroutines_mean"`
	GoroutinesMax      int                `json:"goroutines_max"`
	GCCPUFraction      float64            `json:"gc_cpu_fraction_mean"`
	HeapAllocMeanMB    float64            `json:"heap_alloc_mean_mb"`
	SchedLatencyP99Us  float64            `json:"sched_latency_p99_us"`
	TickCountsTotal    map[string]uint64  `json:"tick_counts_total,omitempty"`
	TickRatesPerSecond map[string]float64 `json:"tick_rates_per_second,omitempty"`
}

// ingestProbe walks energyDir for *.jsonl files (the probe writes one
// per run; harness invocations have exactly one), reads all samples, and
// returns aggregated metrics. Returns (nil, nil) if the directory does
// not exist or contains no JSONL files.
func ingestProbe(energyDir string) (*ProbeAggregates, error) {
	entries, err := os.ReadDir(energyDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var allSamples []energy.Sample
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".jsonl" {
			continue
		}
		samples, err := readJSONLSamples(filepath.Join(energyDir, e.Name()))
		if err != nil {
			return nil, err
		}
		allSamples = append(allSamples, samples...)
	}
	if len(allSamples) == 0 {
		return nil, nil
	}
	sort.Slice(allSamples, func(i, j int) bool {
		return allSamples[i].WallNanos < allSamples[j].WallNanos
	})
	return aggregate(allSamples), nil
}

func readJSONLSamples(path string) ([]energy.Sample, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	var out []energy.Sample
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var s energy.Sample
		if err := json.Unmarshal(line, &s); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, sc.Err()
}

func aggregate(samples []energy.Sample) *ProbeAggregates {
	if len(samples) == 0 {
		return nil
	}
	first := samples[0]
	last := samples[len(samples)-1]

	wallSec := float64(last.WallNanos-first.WallNanos) / nanosPerSecond
	if wallSec < 0 {
		wallSec = 0
	}

	var (
		goroutinesSum int
		goroutinesMax int
		heapSumMB     float64
		schedSumUs    float64
	)
	for _, s := range samples {
		goroutinesSum += s.Goroutines
		if s.Goroutines > goroutinesMax {
			goroutinesMax = s.Goroutines
		}
		heapSumMB += float64(s.HeapAllocBytes) / bytesPerMB
		schedSumUs += s.SchedLatencyP99 * secsToMicros
	}
	n := float64(len(samples))

	gcFrac := 0.0
	if wallSec > 0 {
		gcFrac = (last.GCCPUSeconds - first.GCCPUSeconds) / wallSec
	}

	out := &ProbeAggregates{
		WallSeconds:       wallSec,
		GoroutinesMean:    float64(goroutinesSum) / n,
		GoroutinesMax:     goroutinesMax,
		GCCPUFraction:     gcFrac,
		HeapAllocMeanMB:   heapSumMB / n,
		SchedLatencyP99Us: schedSumUs / n,
	}

	if len(samples) > 1 && wallSec > 0 {
		totals := make(map[string]uint64)
		rates := make(map[string]float64)
		for label, lastN := range last.TickCounts {
			firstN := first.TickCounts[label]
			if lastN < firstN {
				// Counts are monotonic within a single probe lifetime. A
				// decrease here means the probe restarted mid-run (e.g., the
				// lfk subprocess crashed and was relaunched). Drop the label
				// and surface it so the report reader sees the corruption.
				log.Printf("warning: tick count for %q dropped from %d to %d (probe restart?); skipping", label, firstN, lastN)
				continue
			}
			delta := lastN - firstN
			totals[label] = delta
			rates[label] = float64(delta) / wallSec
		}
		if len(totals) > 0 {
			out.TickCountsTotal = totals
			out.TickRatesPerSecond = rates
		}
	}
	return out
}
