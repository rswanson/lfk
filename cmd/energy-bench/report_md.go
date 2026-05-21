package main

import (
	"fmt"
	"sort"
	"strings"
)

// renderMarkdown renders a report as a short markdown document. If baseline
// is non-nil, includes a diff column showing percent change.
func renderMarkdown(r Report, baseline *Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# energy-bench report: %s\n\n", r.Scenario)
	fmt.Fprintf(&b, "fixture=%s contexts=%d duration=%ds reps=%d\n\n",
		r.Fixture, r.Contexts, r.DurationSec, r.Reps)

	renderPrimaryTable(&b, r, baseline)
	renderProbeTable(&b, r, baseline)
	renderTickRateTable(&b, r, baseline)

	if baseline == nil {
		fmt.Fprintln(&b, "\n_(no baseline — run with -update-baseline to create one)_")
	}
	fmt.Fprintln(&b, "\n> Note: fixture-based runs are deterministic but not energetically representative of a real cluster. Use for branch-vs-baseline comparisons only.")
	return b.String()
}

func renderPrimaryTable(b *strings.Builder, r Report, baseline *Report) {
	fmt.Fprint(b, "## Primary metrics (lfk-only via PID-filtered top)\n\n")
	fmt.Fprintln(b, "| metric | current | baseline | delta |")
	fmt.Fprintln(b, "|---|---|---|---|")
	rows := []row{
		{"wakeups_per_s", r.Primary.WakeupsPerSecond, val(baseline, func(r Report) float64 { return r.Primary.WakeupsPerSecond })},
		{"energy_impact", r.Primary.EnergyImpact, val(baseline, func(r Report) float64 { return r.Primary.EnergyImpact })},
		{"pcore_active_residency_pct", r.Primary.PCoreActiveResidency, val(baseline, func(r Report) float64 { return r.Primary.PCoreActiveResidency })},
		{"ecore_active_residency_pct", r.Primary.ECoreActiveResidency, val(baseline, func(r Report) float64 { return r.Primary.ECoreActiveResidency })},
	}
	writeRows(b, rows, baseline != nil)
}

func renderProbeTable(b *strings.Builder, r Report, baseline *Report) {
	if r.Probe == nil {
		fmt.Fprint(b, "\n_(probe metrics unavailable — LFK_ENERGY_PROBE may have been disabled or the JSONL file was empty)_\n")
		return
	}
	fmt.Fprint(b, "\n## Probe metrics (lfk runtime)\n\n")
	fmt.Fprintln(b, "| metric | current | baseline | delta |")
	fmt.Fprintln(b, "|---|---|---|---|")
	rows := []row{
		{"goroutines_mean", r.Probe.GoroutinesMean, valProbe(baseline, func(p *ProbeAggregates) float64 { return p.GoroutinesMean })},
		{"goroutines_max", float64(r.Probe.GoroutinesMax), valProbe(baseline, func(p *ProbeAggregates) float64 { return float64(p.GoroutinesMax) })},
		{"gc_cpu_fraction_mean", r.Probe.GCCPUFraction, valProbe(baseline, func(p *ProbeAggregates) float64 { return p.GCCPUFraction })},
		{"heap_alloc_mean_mb", r.Probe.HeapAllocMeanMB, valProbe(baseline, func(p *ProbeAggregates) float64 { return p.HeapAllocMeanMB })},
		{"sched_latency_p99_us", r.Probe.SchedLatencyP99Us, valProbe(baseline, func(p *ProbeAggregates) float64 { return p.SchedLatencyP99Us })},
	}
	writeRows(b, rows, baseline != nil && baseline.Probe != nil)
}

func renderTickRateTable(b *strings.Builder, r Report, baseline *Report) {
	if r.Probe == nil || len(r.Probe.TickRatesPerSecond) == 0 {
		return
	}
	fmt.Fprint(b, "\n## Tick rates by call site (per second)\n\n")
	fmt.Fprintln(b, "| label | current | baseline | delta |")
	fmt.Fprintln(b, "|---|---|---|---|")
	labels := make([]string, 0, len(r.Probe.TickRatesPerSecond))
	for k := range r.Probe.TickRatesPerSecond {
		labels = append(labels, k)
	}
	sort.Strings(labels)
	for _, label := range labels {
		cur := r.Probe.TickRatesPerSecond[label]
		bl := 0.0
		if baseline != nil && baseline.Probe != nil {
			bl = baseline.Probe.TickRatesPerSecond[label]
		}
		fmt.Fprintf(b, "| %s | %.2f | %.2f | %s |\n", label, cur, bl, deltaPct(cur, bl, baseline != nil && baseline.Probe != nil))
	}
}

type row struct {
	name string
	cur  float64
	bl   float64
}

func writeRows(b *strings.Builder, rows []row, hasBaseline bool) {
	for _, r := range rows {
		fmt.Fprintf(b, "| %s | %.2f | %.2f | %s |\n", r.name, r.cur, r.bl, deltaPct(r.cur, r.bl, hasBaseline))
	}
}

func deltaPct(cur, bl float64, hasBaseline bool) string {
	if !hasBaseline {
		return "—"
	}
	if bl == 0 {
		if cur == 0 {
			return "0%"
		}
		return "n/a (bl=0)"
	}
	pct := (cur - bl) / bl * 100
	return fmt.Sprintf("%+.0f%%", pct)
}

func val(r *Report, f func(Report) float64) float64 {
	if r == nil {
		return 0
	}
	return f(*r)
}

func valProbe(r *Report, f func(*ProbeAggregates) float64) float64 {
	if r == nil || r.Probe == nil {
		return 0
	}
	return f(r.Probe)
}
