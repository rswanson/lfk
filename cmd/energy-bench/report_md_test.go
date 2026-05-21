package main

import (
	"strings"
	"testing"
)

func TestRenderMarkdown_IncludesProbeAndTickRates(t *testing.T) {
	r := Report{
		Scenario:    "idle-foreground",
		Fixture:     "medium",
		Contexts:    3,
		DurationSec: 90,
		Reps:        3,
		Primary: primaryMetrics{
			WakeupsPerSecond: 12.34,
			EnergyImpact:     5.0,
		},
		Probe: &ProbeAggregates{
			GoroutinesMean:    18.5,
			GoroutinesMax:     22,
			GCCPUFraction:     0.012,
			HeapAllocMeanMB:   14.0,
			SchedLatencyP99Us: 320.0,
			TickRatesPerSecond: map[string]float64{
				"pods-refresh":          0.2,
				"kubetris-render-100ms": 0.0,
			},
		},
	}
	out := renderMarkdown(r, nil)
	for _, want := range []string{
		"# energy-bench report: idle-foreground",
		"## Primary metrics (lfk-only via PID-filtered top)",
		"## Probe metrics (lfk runtime)",
		"goroutines_mean",
		"## Tick rates by call site (per second)",
		"pods-refresh",
		"kubetris-render-100ms",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in markdown output:\n%s", want, out)
		}
	}
}

func TestRenderMarkdown_NoProbeSectionWhenAbsent(t *testing.T) {
	r := Report{Scenario: "x", Probe: nil}
	out := renderMarkdown(r, nil)
	if !strings.Contains(out, "probe metrics unavailable") {
		t.Errorf("want 'probe metrics unavailable' note, got:\n%s", out)
	}
}
