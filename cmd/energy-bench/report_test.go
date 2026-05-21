package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReport_PrimaryMetricsAreSetFromInputs(t *testing.T) {
	in := reportInputs{
		Scenario:    "idle-foreground",
		Fixture:     "medium",
		Contexts:    3,
		DurationSec: 90,
		Reps:        3,
		Top: []topSample{
			{CPUPercent: 1.2, PowerScore: 4, WakeupsPerS: 50},
			{CPUPercent: 1.4, PowerScore: 5, WakeupsPerS: 55},
		},
		Power: powermetricsSample{PCoreActiveResidency: 8, ECoreActiveResidency: 22},
	}
	r := buildReport(in)
	assert.InDelta(t, 52.5, r.Primary.WakeupsPerSecond, 0.01)
	assert.InDelta(t, 4.5, r.Primary.EnergyImpact, 0.01)
	assert.Equal(t, 8.0, r.Primary.PCoreActiveResidency)
	assert.Equal(t, 22.0, r.Primary.ECoreActiveResidency)
}

func TestReport_MarkdownIncludesDiffSection(t *testing.T) {
	current := Report{
		Scenario: "idle-foreground",
		Primary:  primaryMetrics{WakeupsPerSecond: 100, EnergyImpact: 5, PCoreActiveResidency: 10},
	}
	baseline := Report{
		Scenario: "idle-foreground",
		Primary:  primaryMetrics{WakeupsPerSecond: 200, EnergyImpact: 10, PCoreActiveResidency: 15},
	}
	md := renderMarkdown(current, &baseline)
	assert.Contains(t, md, "idle-foreground")
	assert.Contains(t, md, "wakeups_per_s")
	assert.Contains(t, md, "-50%") // current is half of baseline
}

func TestReport_RoundTripsJSON(t *testing.T) {
	r := buildReport(reportInputs{Scenario: "active-scripted", Contexts: 3})
	data, err := json.Marshal(r)
	require.NoError(t, err)
	assert.True(t, strings.Contains(string(data), "active-scripted"))
}
