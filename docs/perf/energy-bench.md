# energy-bench: lfk's energy / perf harness

A reproducible, macOS-aware harness for measuring lfk's idle and active
energy behaviour against checked-in baselines.

## Quick start

```
make build
make energy-bench                                  # default: idle-foreground, 90s, 1 rep
SCENARIO=active-scripted make energy-bench         # other scenarios
make energy-bench-all                              # run all three
```

Each run writes `report.json` and `report.md` under
`testdata/energy-reports/<scenario>-<timestamp>/`. The markdown report
compares the current run to the checked-in baseline at
`testdata/baselines/<scenario>.json`.

## Updating a baseline

When an optimisation lands and the new numbers are the new normal:

```
SCENARIO=idle-foreground make energy-bench-update-baseline
git add testdata/baselines/idle-foreground.json
git commit -m "perf: rebaseline idle-foreground after <change>"
```

Baselines are committed; reports are not (they are noise).

## Reading the report

The primary metrics, in order of importance:

1. `wakeups_per_s` — the best proxy for macOS battery drain on Apple
   Silicon. Sourced from `top`'s `IDLEW` (idle wakeups) column.
2. `energy_impact` — what Activity Monitor shows the user. Sourced from
   `top`'s `POWER` column.
3. `pcore_active_residency_pct` / `ecore_active_residency_pct` — work that
   lands on P-cores is ~3-5x more expensive than work on E-cores. Look for
   wins that shift work from P to E even when total CPU% doesn't move.
   Only available when running with `-powermetrics` (sudo).

Supporting diagnostics are rendered inline in `report.md`: goroutine
count, GC CPU fraction, allocation rate, sched latency p99, and
per-callsite tick rates. The raw per-second probe samples are also
preserved under `$LFK_DATA_DIR/energy/<run-id>.jsonl` for ad-hoc
analysis (paths.DataDir() determines the prefix).

## Higher-fidelity runs (sudo)

The default uses `top(1)` (no sudo). To capture package power and accurate
P/E-core residency, use `powermetrics`:

```
sudo -v
go run ./cmd/energy-bench -binary ./lfk -scenario idle-foreground -powermetrics
```

`sudo -v` warms the sudo timestamp so the harness can call `sudo
powermetrics` without prompting mid-run.

## Probe-only mode

The in-process probe also runs in normal lfk sessions:

```
LFK_ENERGY_PROBE=1 ./lfk
```

It writes one JSONL file per session under `$LFK_DATA_DIR/energy/`,
flushed on shutdown and on `SIGUSR1`. Useful for ad-hoc investigation
without the full harness.

## What the harness is not

- It is not a benchmark of hot CPU paths. Use `go test -bench` for that.
- It is not a stand-in for real-cluster validation. The fake apiserver
  emits bookmark events at a fixed cadence, not the messy event stream
  of a production cluster. Use the harness for branch-vs-baseline
  comparisons, not for absolute energy claims.

## What the report contains

The report has three sections:

1. **Primary metrics** — wakeups, energy impact, and (with `-powermetrics`)
   P/E-core active residency. These come from `top(1)` running concurrently
   with the scenario, filtered to lfk's PID. They reflect lfk's specific
   contribution to system energy, not the host's total.

2. **Probe metrics** — goroutines, GC CPU fraction, heap allocation, and
   scheduler latency. These come from the in-process probe and are only
   meaningful for the lfk process. Available when `LFK_ENERGY_PROBE=1` is
   set (the harness sets this automatically).

3. **Tick rates by call site** — the per-second rate at which each
   labelled `energy.Tick` callback fires. This is the most direct signal
   for spotting an offending ticker: a high rate on a label that should
   only fire when its overlay is open means the ticker is leaking.

Note: today the scenario runs exactly once per `energy-bench` invocation
regardless of `-reps`. The flag is preserved as report metadata so that a
future revision can implement multi-rep averaging without breaking the
report schema; for now, treat each run as a single sample. Capture
several runs manually and compare if variance matters for your decision.

## Probe overhead verification

A gating test verifies the in-process probe adds < 1% to `wakeups_per_s`
on its hot path. It is skipped by default (results are noisy on busy
laptops); run it explicitly when validating a probe change:

```
ENERGY_OVERHEAD_TEST=1 go test ./cmd/energy-bench -run TestProbe_OverheadUnder1Percent -v -count=1 -timeout=120s
```

Repeat 3-5 times under quiet system conditions if the first result looks
borderline.

## Pre-run checklist

Variance on a laptop is high. Before a measurement run:

- Plug into power (or consistently *not* plug in, but pick one).
- Quit other CPU-heavy apps.
- Let the machine sit for ~30 seconds after closing apps before starting.
- Run with the same `-reps 1` setting both times (current + baseline).
