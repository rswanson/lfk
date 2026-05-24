package main

import (
	"bufio"
	"strconv"
	"strings"
)

type powermetricsSample struct {
	PCoreActiveResidency float64 // %  averaged across all P-clusters in the sample
	ECoreActiveResidency float64 // %  averaged across all E-clusters in the sample
	PackagePowerMW       float64 // averaged across all "Package Power: N mW" lines, if present
}

// parsePowermetricsOutput extracts P/E cluster active residency and
// package power from a `powermetrics --samplers cpu_power,tasks` text
// dump. Handles two output formats seen across Apple Silicon
// generations:
//
//   - "<X>-Cluster active residency: 12.34%"             (older macOS,
//     single P/E cluster, non-HW summary line).
//   - "<X><N>-Cluster HW active residency: 12.34% (...)" (newer M-series,
//     possibly multiple P clusters labelled P0/P1, only the HW line is
//     emitted).
//
// When both forms exist, the non-HW summary is preferred (matches
// older harness behaviour). When only HW lines exist, those are used
// and any cluster-index variants (P0, P1, ...) are averaged.
//
// PackagePowerMW is not emitted by all Mac models (notably some M4 SoCs);
// it is left at zero when the line is absent.
func parsePowermetricsOutput(s string) (powermetricsSample, error) {
	var (
		out                                        powermetricsSample
		pcoreSum, ecoreSum, pcoreHWSum, ecoreHWSum float64
		pcoreN, ecoreN, pcoreHWN, ecoreHWN, powerN int
		powerSum                                   float64
	)
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		switch {
		case isClusterLine(line, 'E', false):
			ecoreSum += parsePercentAfterResidency(line)
			ecoreN++
		case isClusterLine(line, 'E', true):
			ecoreHWSum += parsePercentAfterResidency(line)
			ecoreHWN++
		case isClusterLine(line, 'P', false):
			pcoreSum += parsePercentAfterResidency(line)
			pcoreN++
		case isClusterLine(line, 'P', true):
			pcoreHWSum += parsePercentAfterResidency(line)
			pcoreHWN++
		case strings.HasPrefix(line, "Package Power:"):
			powerSum += parseMW(line)
			powerN++
		}
	}
	if pcoreN > 0 {
		out.PCoreActiveResidency = pcoreSum / float64(pcoreN)
	} else if pcoreHWN > 0 {
		out.PCoreActiveResidency = pcoreHWSum / float64(pcoreHWN)
	}
	if ecoreN > 0 {
		out.ECoreActiveResidency = ecoreSum / float64(ecoreN)
	} else if ecoreHWN > 0 {
		out.ECoreActiveResidency = ecoreHWSum / float64(ecoreHWN)
	}
	if powerN > 0 {
		out.PackagePowerMW = powerSum / float64(powerN)
	}
	return out, sc.Err()
}

// isClusterLine reports whether line is "<letter>{digits}-Cluster
// [HW] active residency:". wantHW switches between the HW variant
// (newer macOS) and the non-HW summary (older macOS). The two
// variants are mutually exclusive: a HW line never matches as non-HW
// and vice versa.
func isClusterLine(line string, letter byte, wantHW bool) bool {
	if len(line) == 0 || line[0] != letter {
		return false
	}
	rest := line[1:]
	// Strip optional cluster index digits (e.g., the "0" in "P0-Cluster").
	for len(rest) > 0 && rest[0] >= '0' && rest[0] <= '9' {
		rest = rest[1:]
	}
	const sep = "-Cluster "
	if !strings.HasPrefix(rest, sep) {
		return false
	}
	rest = rest[len(sep):]
	if wantHW {
		return strings.HasPrefix(rest, "HW active residency:")
	}
	return strings.HasPrefix(rest, "active residency:")
}

// parsePercentAfterResidency extracts the percent value immediately
// after "active residency:" on a cluster line. Trailing breakdown text
// in parentheses (e.g., per-frequency residency) is discarded.
func parsePercentAfterResidency(line string) float64 {
	const marker = "active residency:"
	_, after, ok := strings.Cut(line, marker)
	if !ok {
		return 0
	}
	rest := strings.TrimSpace(after)
	// Trim everything from the first whitespace or '(' — those mark the
	// end of the value token (or the start of the per-MHz breakdown).
	for j, c := range rest {
		if c == ' ' || c == '\t' || c == '(' {
			rest = rest[:j]
			break
		}
	}
	rest = strings.TrimSuffix(rest, "%")
	f, _ := strconv.ParseFloat(rest, 64)
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
