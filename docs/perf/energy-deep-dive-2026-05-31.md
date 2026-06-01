# Energy and performance deep-dive: lfk fork vs upstream

Date: 2026-05-31
Baseline: upstream `janosmiko/lfk` at merge-base `0e18880` (release v0.13.2 + #317)
Fork tip analysed: `a2d7501` (PR #10) plus the QoS removal in this branch (`perf/remove-qos`)

This document supersedes the QoS findings in
`fork-sync-and-perf-review-2026-05-31.md` (which is kept as a point-in-time
record). QoS routing has since been removed; see section 5.

## 1. Summary

The fork's net behavioural delta over upstream, after removing the energy
probe / `energy-bench` harness (PR #10) and the QoS routing (this branch), is a
focus/idle-aware throttling layer over Bubble Tea's timer-driven work:

| Change | What it does | Verdict |
|---|---|---|
| Watch-tick throttle | 2s -> 30s when blurred or idle >120s | Keep. The real win. |
| Generation-token fix | Guarantees exactly one watch chain | Keep. Protects the throttle from backfiring. |
| Local-output poller throttle | 20Hz -> 2Hz when blurred (PTY/log/capture) | Keep. Narrow but clean. |
| Configurable intervals (#8) | `watch_interval`, `blurred_watch_interval` | Keep. Low-risk flexibility. |
| QoS E-core routing (#7) | (removed) was inert on every shipped binary | Removed in this branch. |

The throttle's clearest, measurable benefit is in **apiserver round-trips and
heavy CPU bursts**, not in the idle wakeup floor -- because a 10 Hz spinner that
both upstream and the fork leave unthrottled pins the CPU awake regardless
(section 6). Gating that spinner is the single biggest remaining idle-energy win.

## 2. Scope and method

"Fork vs upstream" is measured as the diff from the merge-base `0e18880`
(the upstream point the fork last synced to) to the fork tip -- not a naive
`upstream/main..main`, which would mislabel upstream's newer #318 commits as
fork deletions.

This is a static / analytical analysis: every rate below is derived from the
source constants and control flow, cited by file. There is no live benchmark,
because the in-app probe and the `energy-bench` harness that used to measure
this were removed in PR #10. Section 10 gives an external recipe to validate the
predictions with `powermetrics` / `top` / `pprof`.

## 3. The perf-relevant surface

After both removals, the fork changes five things relative to upstream:

1. Watch-tick throttle (focus/blur + 120s foreground-idle) -- `internal/app/commands.go`, `internal/app/update_focus.go`
2. Generation-token watch-tick fix -- `internal/app/messages.go`, `internal/app/update_dispatch_msgs.go`, `internal/app/app.go`
3. Local-output poller throttle -- `pollInterval()` in `commands.go`, applied in `ptyexec.go`, `update_capture.go`, `update_exec_log_msgs.go`
4. Configurable intervals (#8) -- `internal/ui/config.go`, `internal/app/options.go`, CLI flags in `main.go`
5. QoS E-core routing (#7) -- removed in this branch (was `internal/qos`).

## 4. Mechanism analysis

### 4.1 Watch-tick throttle (the headline win)

Constants (verified): `DefaultWatchInterval = 2s`, `DefaultBlurredWatchInterval
= 30s`, `foregroundIdleThreshold = 120s` (`commands.go`, `config.go`).
`activeWatchInterval()` returns the blurred value when `!focused ||
foregroundIdle()`, else the foreground value. Each tick runs
`refreshCurrentLevel()` -- a network LIST against the apiserver (x N in union
mode, N <= 8), JSON decode, and a full table re-render.

| State | Upstream cadence | Fork cadence | Network LISTs/hr |
|---|---|---|---|
| Foreground, active | 2s | 2s (unchanged) | 1800 -> 1800 |
| Window blurred | 2s | 30s | 1800 -> 120 |
| Foreground idle >120s | 2s | 30s | 1800 -> 120 |

Upstream LISTs the cluster every 2s forever -- backgrounded on a second
monitor, or sitting idle overnight, it is still doing 1800 apiserver
round-trips/hour and 1800 full re-renders/hour. The fork cuts that 15x to
120/hr whenever the window is unfocused or untouched for 2 minutes. On a large
cluster each LIST is a TLS round-trip plus a decode of potentially thousands of
objects, so this saves real CPU, network, and apiserver load.

UX cost: essentially none. On return, `updateFocus` (focus regain) and
`snapBackIfIdle` (first key/mouse after idle) both fire an immediate
`refreshCurrentLevel()` and restart a fast chain, so the user never sees stale
data -- a fresh fetch lands the instant they come back, then 2s cadence resumes.

### 4.2 Generation-token fix (load-bearing correctness)

`tea.Tick` is non-cancellable, so every focus regain / idle snap-back /
watch-mode toggle used to spawn a parallel self-re-arming chain. Flap focus a
few times and you would be running 2, 3, 4 chains at once -- more refreshes than
upstream, the opposite of the goal.

The fix (`commands.go` `nextWatchTick`, `update_dispatch_msgs.go:129`,
`update_focus.go`, `update_keys.go`, `app_init_cmd.go`): a `watchTickGen` epoch
stamped into `watchTickMsg{gen}`. `nextWatchTick()` bumps the epoch at every
chain-start site; `updateWatchTick` drops any tick whose gen no longer matches
and re-arms only the live chain. Guarantees exactly one chain. Covered by
`watch_tick_gen_test.go` and `update_focus_test.go`.

This is the most important item in the set: without it, the throttle feature
itself could regress into a refresh amplifier.

### 4.3 Local-output poller throttle

`pollInterval(nominal)` floors blurred polling at 500ms (`blurredPollFloor`),
keyed on `m.focused` only -- not idle. That is correct: a focused terminal with
scrolling output is being watched even without keypresses, so only losing focus
should slow it.

| Poller | Focused | Blurred | Reduction |
|---|---|---|---|
| Embedded PTY output (`ptyexec.go`) | 50ms (20Hz) | 500ms (2Hz) | 10x |
| Exec-log refresh (`update_exec_log_msgs.go`) | 50ms (20Hz) | 500ms (2Hz) | 10x |
| Live capture (`update_capture.go`) | 100ms (10Hz) | 500ms (2Hz) | 5x |

The 20Hz PTY poll is the highest-frequency timer in the app, so backing it to
2Hz when you tab away from an embedded shell is the largest per-timer wakeup cut
available. Active only in exec/log/capture modes.

### 4.4 Configurable intervals (#8)

`watch_interval` + `blurred_watch_interval` (config keys plus
`--watch-interval` / `--blurred-watch-interval` flags), clamped to
[500ms, 10m]. Dashboard-on-a-second-monitor users lower the blurred value for
freshness; battery-conscious users raise it.

## 5. QoS E-core routing (#7): removed

Removed in this branch. Rationale, confirmed from source:

- `.goreleaser.yaml` sets `CGO_ENABLED=0` for the release build.
- The QoS syscall lived in `qos_darwin.go` behind `//go:build darwin && cgo`,
  with `qos_other.go` providing `const qosSupported = false` for every other
  build. `RunWith` only locked a thread and set the QoS class when
  `qosSupported && LFK_QOS != "off"`.

So on every release / Homebrew / CI binary, `qos.RunWith` was a pure no-op: it
just called `fn()`. The informer goroutine was never E-core-routed. It activated
only for users who did a local `go build` / `go install` on macOS with a C
compiler present -- and even then the payoff was speculative, because the
informer watch loop is I/O-blocked on the watch stream (already cheap) and a
Utility QoS hint only matters under CPU contention.

The earlier follow-up (gating `LockOSThread` on `qosSupported`) had already
removed the *cost* -- the old code stranded one locked OS thread per informer on
no-cgo builds for zero benefit. So the state before this change was "no cost,
~zero benefit on shipped binaries." Removing the package entirely eliminates the
dead code and the cgo build-tag complexity, and reduces upstream drift.

The change: delete `internal/qos/`, drop the import, and revert the informer
goroutine in `informer_cache.go` to a plain goroutine. (The unrelated
Kubernetes Pod QoS-class column and crash-investigator fields are untouched.)

## 6. The idle-wakeup ceiling: an unthrottled 10 Hz spinner

The dominant idle wakeup source in lfk is not the watch tick -- it is the
spinner, and neither upstream nor the fork throttles it.

Verified: `Init` (`app_init.go`) unconditionally seeds `m.spinner.Tick`, and
`updateTick` (`update.go`) re-arms it on every `spinner.TickMsg` with no
`m.loading` gate. The bubbles `spinner.Dot` default is `FPS = time.Second/10`,
so this is a ~10 Hz timer for the entire lifetime of the process, foreground or
background, idle or busy, loading or not. Each tick is a frame increment plus a
full `View()` composition pass (Bubble Tea re-renders after every Update; the
cell-diff suppresses terminal writes, but the layout recomposition CPU is spent
regardless).

Why this matters for energy specifically: deep package C-states require the CPU
to stay asleep for long stretches. A 10 Hz timer wakes it every 100ms no matter
what, so even with the watch tick throttled to 30s the CPU never sleeps deeper
than 100ms. That caps how much battery the watch-tick throttle can save. The
throttle's clearest wins are the network round-trips and the heavy LIST+render
CPU bursts it eliminates -- not the idle-wakeup floor, which the spinner pins.

This is shared upstream behaviour, not a fork regression. But it is the single
biggest idle-energy win still available (section 9).

## 7. Correctness and portability caveats

- Focus path needs a focus-reporting terminal. `main.go` sets
  `tea.WithReportFocus()`, so lfk requests DECSET-1004. Terminals that do not
  implement it -- Terminal.app, plain `ssh`, bare `tmux` without
  `focus-events on` -- silently ignore the request, so `FocusMsg`/`BlurMsg`
  never arrive and the blur half of the throttle is inert there. The 120s
  foreground-idle path is pure `time.Since(lastInputAt)` and works on every
  terminal, so it is the portable, guaranteed win; the blur path is a bonus on
  iTerm2 / kitty / ghostty / wezterm / alacritty / tmux-with-focus-events.

- Idle timer is robust against a passive mouse. `main.go` uses
  `tea.WithMouseCellMotion()` (not `AllMotion`), which reports motion only
  during a drag. A stationary mouse hovering over the window does not generate a
  `MouseMsg` and does not reset the idle clock, so idle throttling triggers as
  intended. Only real keys/clicks/drags/scroll reset it (`update.go`).

- No energy regressions. The added per-tick gen compare and per-keystroke
  `snapBackIfIdle()` time check are negligible, and the fork adds no new
  always-on goroutines or timers.

## 8. Net assessment

The watch-tick throttle plus the generation-token fix are genuinely good: a real
15x cut in apiserver LISTs and re-renders when idle/backgrounded, no UX cost,
and the gen-token fix prevents the feature from backfiring. The poller throttle
is a clean, well-scoped 10x cut for embedded shells. Configurable intervals add
low-risk flexibility. QoS #7 added no value on shipped binaries and is now gone.

## 9. Recommendations (ranked)

1. Gate the spinner tick on `m.loading`. Biggest remaining energy win, and it
   helps foreground too: start the chain when a load begins, stop re-arming when
   nothing is loading. Cuts the idle wakeup floor from ~10 Hz to ~0, letting the
   CPU reach deep sleep -- which is what finally lets the watch-tick throttle's
   idle savings materialise. Worth upstreaming.
2. Document the focus-reporting dependency in `docs/usage.md`: blur throttling
   needs a DECSET-1004 terminal; the 120s idle throttle works everywhere.
3. Note the always-on informer cost. When the informer cache is promoted
   (`auto` mode, lists >= 1000 items), it keeps a watch stream open and indexes
   every object change continuously, regardless of focus -- background
   CPU/network the watch-tick throttle does not touch. On a churny cluster this
   can dominate idle energy more than the watch tick. Measure before optimising.
4. Restore a measurement path. Since `energy-bench` is gone, keep section 10's
   external recipe in `docs/perf/` so these claims stay falsifiable.

## 10. How to measure it for real

Against a live (or kind/minikube) cluster, foreground vs backgrounded:

```
# Idle wakeups + CPU for the lfk process, 5s samples:
top -pid $(pgrep -x lfk) -stats pid,cpu,idlew -l 0 -s 5

# System-wide energy attribution (needs sudo); watch the lfk row:
sudo powermetrics --samplers tasks --show-process-energy -i 5000

# CPU profile of an idle session (pprof is already wired upstream):
LFK_PPROF_ADDR=127.0.0.1:6060 lfk &
go tool pprof -seconds=30 http://127.0.0.1:6060/debug/pprof/profile
```

Prediction this analysis makes: foreground idle-wakeups are dominated by the
~10 Hz spinner (~10-12/s) and barely move with the watch throttle; the
throttle's signal shows up as a ~15x drop in apiserver LIST traffic and in the
periodic CPU spikes from `refreshCurrentLevel`, not in the wakeup floor -- until
the spinner is gated, after which the idle wakeup floor collapses.
