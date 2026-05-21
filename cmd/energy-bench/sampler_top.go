package main

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// topSample holds one row from `top -l N -stats pid,cpu,csw` output.
// On this macOS version the available stats are pid, cpu, and csw (context
// switches). The CSW count is stored in WakeupsPerS as a best available proxy
// for wakeup activity; it is a cumulative counter, not a per-second rate.
type topSample struct {
	PID         int
	CPUPercent  float64
	WakeupsPerS float64 // populated from the CSW (context-switch) column
}

// parseTopOutput parses `top -l N -stats pid,cpu,csw` output and returns one
// slice element per data row. Header and summary lines are skipped. Rows that
// cannot be parsed are silently skipped so that a single bad line does not
// abort the whole parse.
//
// macOS `top -l` (logging mode) repeats the header block for every sample
// interval. Each block begins with summary lines (Processes, Load Avg, …)
// followed by a column-header line that starts with "PID", then data rows.
// The parser re-arms on every "PID" header so multiple intervals work.
//
// CSW values may be suffixed with '+' (indicating counter wraparound);
// the suffix is stripped before parsing.
func parseTopOutput(s string) ([]topSample, error) {
	if strings.TrimSpace(s) == "" {
		return nil, errors.New("empty top output")
	}

	var out []topSample
	sc := bufio.NewScanner(strings.NewReader(s))
	headerSeen := false

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// A line starting with "PID" is the column header — reset state so
		// that the immediately following lines are treated as data rows.
		if strings.HasPrefix(line, "PID") {
			headerSeen = true
			continue
		}

		if !headerSeen {
			continue
		}

		fields := strings.Fields(line)
		// Expect at least: PID %CPU CSW
		if len(fields) < 3 {
			continue
		}

		pid, err := strconv.Atoi(fields[0])
		if err != nil {
			// Not a numeric PID — probably a summary or separator line;
			// treat it as a section boundary and wait for the next header.
			headerSeen = false
			continue
		}

		cpu := parseFloatLoose(fields[1])
		csw := parseFloatLoose(fields[2])

		out = append(out, topSample{
			PID:         pid,
			CPUPercent:  cpu,
			WakeupsPerS: csw,
		})
	}

	return out, sc.Err()
}

// parseTopOutputForPID returns only the rows from the parsed output whose PID
// field matches the given pid.
func parseTopOutputForPID(s string, pid int) ([]topSample, error) {
	all, err := parseTopOutput(s)
	if err != nil {
		return nil, err
	}

	out := make([]topSample, 0, len(all))
	for _, x := range all {
		if x.PID == pid {
			out = append(out, x)
		}
	}

	return out, nil
}

// parseFloatLoose strips common non-numeric suffixes ('%', '+') and parses the
// remainder as a float64. Unparseable values return 0.
func parseFloatLoose(s string) float64 {
	s = strings.TrimSuffix(s, "%")
	s = strings.TrimSuffix(s, "+")
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
