package main

import (
	"fmt"
	"strings"
)

// renderMarkdown renders a report as a short markdown document. If baseline
// is non-nil, includes a diff column showing percent change.
func renderMarkdown(r Report, baseline *Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# energy-bench report: %s\n\n", r.Scenario)
	fmt.Fprintf(&b, "fixture=%s contexts=%d duration=%ds reps=%d\n\n",
		r.Fixture, r.Contexts, r.DurationSec, r.Reps)
	fmt.Fprintln(&b, "| metric | current | baseline | delta |")
	fmt.Fprintln(&b, "|---|---|---|---|")
	rows := []struct {
		name    string
		cur, bl float64
	}{
		{"wakeups_per_s", r.Primary.WakeupsPerSecond, val(baseline, func(r Report) float64 { return r.Primary.WakeupsPerSecond })},
		{"energy_impact", r.Primary.EnergyImpact, val(baseline, func(r Report) float64 { return r.Primary.EnergyImpact })},
		{"pcore_active_residency_pct", r.Primary.PCoreActiveResidency, val(baseline, func(r Report) float64 { return r.Primary.PCoreActiveResidency })},
		{"ecore_active_residency_pct", r.Primary.ECoreActiveResidency, val(baseline, func(r Report) float64 { return r.Primary.ECoreActiveResidency })},
	}
	for _, row := range rows {
		delta := "—"
		if baseline != nil && row.bl != 0 {
			pct := (row.cur - row.bl) / row.bl * 100
			delta = fmt.Sprintf("%+.0f%%", pct)
		}
		fmt.Fprintf(&b, "| %s | %.2f | %.2f | %s |\n", row.name, row.cur, row.bl, delta)
	}
	if baseline == nil {
		fmt.Fprintln(&b, "\n_(no baseline — run with --update-baseline to create one)_")
	}
	fmt.Fprintln(&b, "\n> Note: fixture-based runs are deterministic but not energetically representative of a real cluster. Use for branch-vs-baseline comparisons only.")
	return b.String()
}

func val(r *Report, f func(Report) float64) float64 {
	if r == nil {
		return 0
	}
	return f(*r)
}
