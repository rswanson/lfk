# Energy / battery results

Tracks the measured impact of lfk's energy work. The harness, scenarios,
and metric definitions live in [energy-bench.md](./energy-bench.md); this
file is the running scorecard.

`wakeups_per_s` (macOS `top` IDLEW) is the headline metric — the best proxy
for Apple-Silicon battery drain. Lower is better.

## Optimizations landed

| PR | Change | Scenario it moves | Mechanism |
|----|--------|-------------------|-----------|
| #4 | Throttle watch tick when terminal is unfocused | idle-background | watch tick 2s -> blurred interval (default 30s) on `tea.BlurMsg` |
| #5 | Slow watch tick when foreground-idle 120s | idle-foreground | same blurred cadence after 120s of no key/mouse input |
| #7 | Route informer + probe goroutines to E-cores | all (energy, not wakeups) | macOS QoS `Utility`/`Background` via `qos.RunWith` |
| (this branch) | Configurable blurred/idle interval | idle-* | `blurred_watch_interval` config + `--blurred-watch-interval` |
| (this branch) | Throttle local output pollers when blurred | active PTY/log/capture | `pollInterval()` floors PTY/exec-log/capture to 500ms while unfocused |

## Analytical wakeup reduction (first-principles, no run required)

These follow directly from the tick cadences and hold regardless of host:

| Ticker | Focused/active | Blurred or idle | Reduction |
|--------|----------------|-----------------|-----------|
| `main-context-interval` (watch) | 2s = 0.50/s | 30s = 0.033/s | ~15x |
| `ptyexec-50ms` / `exec-log-50ms` | 50ms = 20/s | 500ms = 2/s | 10x |
| `capture-100ms` | 100ms = 10/s | 500ms = 2/s | 5x |

Before PR #4 there was no throttle at all: the watch tick ran at the
foreground cadence even when blurred, and the pollers still do today on
`main` (this branch is the first to back them off).

## Empirical results (fill from a macOS run)

The harness is macOS-only (`top` IDLEW, `powermetrics`, QoS). Capture on
Apple Silicon, ideally with `sudo -v` first for P/E-core residency:

```
make build
sudo -v && SCENARIO=idle-background    go run ./cmd/energy-bench -binary ./lfk -scenario idle-background    -powermetrics
sudo -v && SCENARIO=idle-foreground    go run ./cmd/energy-bench -binary ./lfk -scenario idle-foreground    -powermetrics
sudo -v && SCENARIO=active-scripted     go run ./cmd/energy-bench -binary ./lfk -scenario active-scripted     -powermetrics
```

| Scenario | Baseline `wakeups_per_s` | Current `wakeups_per_s` | Delta |
|----------|--------------------------|-------------------------|-------|
| idle-background | 2.62 (pre-#4) | _TODO_ | _TODO_ |
| idle-foreground | 2.77 (pre-#4) | _TODO_ | _TODO_ |
| active-scripted | 3.00 | _TODO_ | _TODO_ |

> Caveat: the checked-in baselines under `testdata/baselines/` were last
> captured at PR #3 and never refreshed, so they predate the #4/#5/#7
> wins. A run today therefore measures the cumulative gain. Once recorded,
> re-baseline so future diffs start from the new normal:
>
> ```
> SCENARIO=idle-background make energy-bench-update-baseline
> ```

## Open items

- `status-clear-5s` fires ~1/s in the idle baselines even though the
  automatic watch-refresh paths do not arm it. Trace the exact call site
  with the JSONL probe (`LFK_ENERGY_PROBE=1 ./lfk`, then inspect
  `$LFK_DATA_DIR/energy/<run-id>.jsonl`) to confirm the source before
  optimizing.
