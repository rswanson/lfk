package main

import "math"

type primaryMetrics struct {
	WakeupsPerSecond     float64 `json:"wakeups_per_s"`
	EnergyImpact         float64 `json:"energy_impact"`
	PCoreActiveResidency float64 `json:"pcore_active_residency_pct"`
	ECoreActiveResidency float64 `json:"ecore_active_residency_pct"`
}

type Report struct {
	Scenario    string         `json:"scenario"`
	Fixture     string         `json:"fixture"`
	Contexts    int            `json:"contexts"`
	DurationSec int            `json:"duration_s"`
	Reps        int            `json:"reps"`
	Primary     primaryMetrics `json:"primary"`
}

type reportInputs struct {
	Scenario    string
	Fixture     string
	Contexts    int
	DurationSec int
	Reps        int
	Top         []topSample
	Power       powermetricsSample
}

func buildReport(in reportInputs) Report {
	wakeups, energy := 0.0, 0.0
	if len(in.Top) > 0 {
		for _, t := range in.Top {
			wakeups += t.WakeupsPerS
			energy += t.PowerScore
		}
		wakeups /= float64(len(in.Top))
		energy /= float64(len(in.Top))
	}
	return Report{
		Scenario:    in.Scenario,
		Fixture:     in.Fixture,
		Contexts:    in.Contexts,
		DurationSec: in.DurationSec,
		Reps:        in.Reps,
		Primary: primaryMetrics{
			WakeupsPerSecond:     round2(wakeups),
			EnergyImpact:         round2(energy),
			PCoreActiveResidency: in.Power.PCoreActiveResidency,
			ECoreActiveResidency: in.Power.ECoreActiveResidency,
		},
	}
}

func round2(f float64) float64 { return math.Round(f*100) / 100 }
