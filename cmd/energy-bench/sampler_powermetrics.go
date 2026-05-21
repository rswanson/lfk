package main

import (
	"bufio"
	"strconv"
	"strings"
)

type powermetricsSample struct {
	PCoreActiveResidency float64 // %
	ECoreActiveResidency float64 // %
	PackagePowerMW       float64 // averaged across all "Package Power: N mW" lines
}

// parsePowermetricsOutput extracts the three fields the harness primary
// metrics need from a `powermetrics --samplers cpu_power,tasks` text dump.
// macOS labels look like "P-Cluster active residency: 12.34%" and
// "Package Power: 567 mW". If Package Power appears multiple times (once
// per sample interval), the values are averaged.
func parsePowermetricsOutput(s string) (powermetricsSample, error) {
	var out powermetricsSample
	var powerSum float64
	var powerN int
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		switch {
		case strings.HasPrefix(line, "P-Cluster active residency:"):
			out.PCoreActiveResidency = parsePercent(line)
		case strings.HasPrefix(line, "E-Cluster active residency:"):
			out.ECoreActiveResidency = parsePercent(line)
		case strings.HasPrefix(line, "Package Power:"):
			powerSum += parseMW(line)
			powerN++
		}
	}
	if powerN > 0 {
		out.PackagePowerMW = powerSum / float64(powerN)
	}
	return out, sc.Err()
}

func parsePercent(line string) float64 {
	i := strings.LastIndex(line, ":")
	if i < 0 {
		return 0
	}
	v := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(line[i+1:]), "%"))
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func parseMW(line string) float64 {
	i := strings.LastIndex(line, ":")
	if i < 0 {
		return 0
	}
	v := strings.TrimSpace(line[i+1:])
	v = strings.TrimSuffix(v, "mW")
	v = strings.TrimSpace(v)
	f, _ := strconv.ParseFloat(v, 64)
	return f
}
