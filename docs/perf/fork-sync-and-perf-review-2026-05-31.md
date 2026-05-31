# Fork sync + performance review (2026-05-31)

This report covers two things the maintainer asked for:

1. **Sync** of `rswanson/lfk` with upstream `janosmiko/lfk` (pull upstream
   commits in, resolve conflicts).
2. An **aggressive review** of the fork's performance/energy changes against
   the new upstream code — do they actually improve performance without hurting
   the user experience?

Method: the sync was done as a real merge in an isolated worktree
(`sync/upstream-v0.13.2`); the review was run as a 6-dimension adversarial
code audit (one critical reviewer per area, each high/critical finding then
re-checked by an independent verifier instructed to *refute* it). Everything
below is static-analysis + the merged tree's own test/lint/build results. No
on-cluster energy measurement was run (see "What we could not verify").

---

## 1. What was done — the sync

| | |
|---|---|
| Merge-base | `2a64375` (release 0.12.3) |
| Upstream pulled in | 36 commits → `2e3c0d9` (releases 0.12.4 → **0.13.2**) |
| Fork commits preserved | 32 (energy-bench harness + 4 perf PRs) |
| Conflicts | **1 file** — `main.go` (imports only) |
| Result | merge commit `6e94a4e`, branch `sync/upstream-v0.13.2`, **0 behind upstream** |

### Upstream features now merged in
Security findings dashboard (Trivy/Falco/Gatekeeper/Kubescape/PolicyReport),
k9s-style views/columns config + column registry, pin individual resource
types, cluster dashboard rendering fixes, per-kind sort memory + per-context
column memory, negative namespace selection, CRD printer-column/display-name
fixes, and a **graceful-shutdown force-quit watchdog** (10s) — plus dependency
bumps.

### The one conflict and how it was resolved
`main.go` imports only. Both sides edited the same import block; the function
bodies merged cleanly. Resolution combined both sides:
- kept the fork's `os/signal` + `syscall` (SIGUSR1 energy-probe flush) and the
  `internal/perf/energy` import;
- kept the fork's `tea.WithReportFocus()` program option (focus/blur throttle);
- kept upstream's `time` import (force-quit watchdog `armForceQuit`);
- **dropped** the now-unused `internal/model` import — upstream removed its only
  consumer (`model.PinnedGroups`), so keeping it would not compile.

### Verification of the merged tree (all green)
- `go build ./...` — ok
- `go vet ./...` — ok
- `golangci-lint run ./...` — **0 issues** (also enforced by the pre-commit hook on the merge commit)
- `go test ./...` — **23 packages ok, 0 failures** (incl. `internal/qos`,
  `internal/perf/energy`, `cmd/energy-bench`, `internal/app`)
- Merge integrity spot-check: the fork's focus/idle throttle is still correctly
  wired after the auto-merge — `update.go` dispatches `Blur/FocusMsg` and resets
  `lastInputAt` + calls `snapBackIfIdle` on key/mouse; `updateWatchTick` still
  re-arms via `activeWatchInterval()`. Nothing upstream overrode the cadence.

> The merge lives on branch `sync/upstream-v0.13.2`, **not** `main`.
> Fast-forwarding `main` and pushing was intentionally left to the maintainer.

---

## 2. The performance changes under review

Four perf PRs plus the measurement harness:

| PR | Commit | What it does |
|----|--------|--------------|
| #3 | `5dc4d56` | In-process energy probe; flush JSONL on PTY teardown / SIGUSR1 |
| #4 | `918fdcf` | Throttle watch tick to 30s when the terminal is **unfocused** |
| #5 | `ff39527` | Slow watch tick to 30s after **120s of no input** (foreground-idle) |
| #7 | `b05c985` | Route informer + probe goroutines to **E-cores** via macOS QoS |
| — | (many) | `cmd/energy-bench`: PTY-driven energy/perf measurement harness |

### Per-area verdict

| Area | Verdict |
|------|---------|
| #3 Energy probe | ✅ **Worked (mostly)** — off-by-default, ~free when off, race-safe, really wired in |
| #4 Focus/blur throttle | ⚠️ **Partially** — logic correct, but a duplicate-tick-chain bug + silent no-op on common terminals |
| #5 Foreground-idle slowdown | ⚠️ **Partially** — logic correct & well-tested, but the same duplicate-tick-chain bug + no freshness guard |
| #7 QoS E-core routing | ❌ **Did not (in production)** — inert in shipped binaries (`CGO_ENABLED=0`) |
| energy-bench harness | ⚠️ **Partially** — real and runnable, but weak rigor (1 run, no variance, no residency data) |
| Merge integrity | ✅ **Worked (with caveats)** — coherent; shutdown drain is deadlock-safe |

---

## 3. What worked

- **The energy probe (#3) is the strongest piece.** Verified: off unless
  `LFK_ENERGY_PROBE=1`; no goroutine spawned when off; the hot path
  (`energy.Tick`, called from ~24 sites) is a lock-free `atomic.Pointer` load +
  nil-safe bool check, no allocation in the disabled callback. Race-safe: ring
  buffer under `ringMu`, tick counts under `tickMu`, lifecycle under `mu`,
  `flush()` serialized by a package-level `flushMu` so the SIGUSR1 handler can't
  race `Stop()`. `Stop()` is synchronous (`close(stopCh)` → `wg.Wait()`) and
  idempotent. A migration test even asserts every `tea.Tick` in `internal/app`
  routes through `energy.Tick`, and there's an adversarial <1%-overhead gate.
- **The focus/idle *logic* is correct and well-tested.** Idle is wall-clock
  based (`time.Since(lastInputAt)`), so no tick drift; key, mouse, and focus
  events all reset the clock; focus/idle compose cleanly (blur wins; snap-back
  defers to the blur path when blurred). On focus regain there's an immediate
  refresh, so the user never stares at 30s-stale data after tabbing back.
- **The merge is coherent and shutdown is safe.** `signalShutdown()` cancels
  before `drainShutdown()` waits, so the drain cannot deadlock against the
  QoS-locked informer goroutines (they exit on `stopCh` close, releasing the
  locked OS thread) or the probe. No lock-ordering cycle.
- **QoS is implemented the *right* way for what it tries to do.** `qos.RunWith`
  uses `runtime.LockOSThread()` and never unlocks — which correctly sidesteps
  the classic Go pitfall (per-thread QoS leaking to other goroutines via thread
  reuse). The class constants (`Utility=17`, `Background=9`) match `<sys/qos.h>`.
  Choosing `Utility` (not `Background`) for the informer is the right call to
  bound UX lag. **It just doesn't run in shipped builds** — see below.
- **The energy-bench harness exists, compiles, and is wired end-to-end** with
  *real* committed baselines (not placeholders) and an honest
  `docs/perf/energy-bench.md` that states its own limits.

---

## 4. What didn't work (or doesn't hold up)

### 🔴 Highest-impact: the watch-tick "duplicate chain" bug (PR #4 **and** #5, same root cause)
Adversarially verified, not refuted (verifier downgraded severity high→medium
and corrected the category from "race" — Bubble Tea is single-goroutine, so
it's a *scheduling/resource-leak* bug, not a data race).

- `watchTickMsg` carries **no generation token**, and `scheduleWatchTick` is
  dispatched from four places: `Init`, the self-re-arming `updateWatchTick`
  chain, `updateFocus` (every FocusMsg), and the `r` key.
- `tea.Tick` is one-shot and **non-cancellable**, and `energy.Tick` does no
  label dedup. So when focus is regained — or when idle ends via `snapBackIfIdle`
  — a **new** self-perpetuating tick chain starts while the old one is still
  re-arming. Result: two (or more, across repeated flaps) concurrent watch
  chains, each calling `refreshCurrentLevel` forever.
- **This inverts the energy goal**: under exactly the focus-churn / idle-exit
  these PRs target, refresh frequency, API load, CPU, and wakeups *increase*.
- **The fix already exists in the same package**: `previewDebounceTickMsg{gen}`
  guards stale ticks (`if msg.gen != m.previewDebounceGen { return }`). Mirror
  it: add `m.watchTickGen`, bump it in `updateFocus` / `updateBlur` / `snapBackIfIdle`
  / `r`-key, stamp it into `watchTickMsg`, and drop stale ticks in
  `updateWatchTick`. Add a `blur→focus→blur→focus` regression test asserting one
  surviving chain.

### 🔴 QoS E-core routing (#7) is inert in every released binary
Confirmed directly: `.goreleaser.yaml:11` sets `CGO_ENABLED=0` for the only
build (linux/darwin/windows). `qos_darwin.go` is `//go:build darwin && cgo`, so
the real `pthread_set_qos_class_self_np` is compiled **out** and the no-op
`qos_other.go` is linked in. Every distributed artifact (Homebrew, archives,
nfpm, docker, …) therefore does **zero** E-core routing — while still paying
`runtime.LockOSThread()` per informer (one permanently-locked OS thread per
`(context, GVR)`, potentially dozens under auto/always cache mode). The headline
optimization only runs in a local `cgo` macOS build. There is also no opt-out
(`LFK_QOS=off`), no battery/AC awareness, and the cgo-only tests give false
confidence by exercising a path that's never shipped.

### 🟠 Additional finding (found directly, not by the workflow): idle/blur can *speed up* a slow configured interval
`activeWatchInterval()` returns `blurredWatchInterval` (30s) **unconditionally**
when blurred or idle. But `watch-interval` is configurable up to
`MaxWatchInterval = 10m`. A user who sets `watch-interval: 5m` would see refreshes
*accelerate* to 30s the moment they look away — the opposite of the intent.
**Fix:** `return max(m.watchInterval, blurredWatchInterval)` so the throttle only
ever slows down, never speeds up.

### 🟠 The optimizations are not backed by data
- The committed baselines are **single runs** (`reps: 1`); the `-reps` flag is
  advisory and silently ignored. No warmup gating, no std-dev/CI/significance.
  For sub-percent idle-tick and P-vs-E deltas, that's below a laptop's noise floor.
- **P/E-core residency is `0` in every committed baseline** (powermetrics needs
  sudo, which the default harness doesn't use) — so the QoS-to-E-core change has
  *no* residency data behind it, and `0.0` reads misleadingly like a real
  measurement rather than "not measured."
- The committed `active-scripted` baseline is internally suspicious as a guard:
  `preview-debounce` fires at **10.39/s** and `main-refresh-5s` at **3.18/s**
  (nominal 0.2/s) — either expected-under-typing-but-undocumented, or a ticker
  leak the harness *recorded but did not flag*.

### 🟡 UX gaps
- **`WithReportFocus` is a silent no-op on common terminals** (macOS
  Terminal.app, plain ssh, tmux/screen without `focus-events on`). There it's
  harmless (private DECSET, no byte leakage) but the focus throttle and the
  snap-back refresh simply never engage — only the portable input-idle path works.
- **No freshness guard during critical events.** A user watching a
  CrashLoopBackOff/rollout who stops typing for 120s drops to 30s refresh
  exactly when fresh data matters most. No "important activity" exemption.
- **Focus refresh isn't gated by overlays/exec/PTY** — regaining focus while a
  confirm dialog or exec session is open still fires a background list reload.

### 🟡 Process / packaging
- The dev-only probe library is linked into the **production** binary (small, but
  adds an env-gated file writer + SIGUSR1 meaning). Consider a `//go:build`
  tag with a no-op release stub.
- Final flush is skipped on `os.Exit` paths (force-quit watchdog `os.Exit(0)`,
  root-error `os.Exit(1)`) and there's no periodic flush, so a hard exit loses
  the whole in-memory run. Flush in the watchdog closure and/or add a periodic flush.

---

## 5. Recommended changes (priority order)

1. **Fix the duplicate-tick-chain bug** (one fix resolves both PR #4 and #5).
   Add a `watchTickGen` epoch to `watchTickMsg`, mirroring the existing
   `previewDebounceTickMsg{gen}` idiom; drop stale ticks in `updateWatchTick`;
   add a flap regression test. *This is the single highest-value change — until
   it lands, the throttle PRs can make energy use worse, not better.*
2. **Clamp the throttle to never speed up**: `max(m.watchInterval, blurredWatchInterval)`.
3. **Decide PR #7's fate.** Either ship a `cgo`-enabled darwin/arm64 release
   build so QoS actually runs (accepting the static-binary/scorecard tradeoffs),
   **or** drop/clearly-document it as source-build-only and remove the
   unconditional `LockOSThread` cost from no-cgo builds. Add an `LFK_QOS=off`
   escape hatch and a startup debug-log line when QoS is actually applied.
4. **Make the harness honest about confidence**: implement the `-reps` loop with
   mean ± stddev (or make `-reps>1` an error until then); mark unmeasured P/E
   residency as "not measured" rather than `0.0`; capture at least one sudo
   `-powermetrics` baseline to back PR #7; investigate the anomalous
   `active-scripted` tick rates.
5. **Close the UX gaps**: gate the focus refresh on `watchMode` + no open
   overlay/exec; add a "non-steady-state resources visible" exemption (or config
   flag) to the idle slowdown; document that focus-based throttling needs
   DECSET-1004 support.
6. **Tidy packaging**: build-tag the probe out of release binaries; flush on the
   `os.Exit` paths and/or periodically.

---

## 6. What we could not verify (limits of this review)

- **No real energy measurement** was run here — this is static analysis plus the
  fork's own (single-run, no-residency) baselines. The actual P-vs-E and
  idle-vs-active wins remain *claimed*, not *measured*. They should be re-measured
  on a `cgo` build **after** the duplicate-chain fix (which currently inflates
  post-idle-exit work).
- Whether the watch tick issues **real API requests** vs only local
  informer-cache reads + repaints wasn't fully traced; the energy attribution
  (fewer requests vs fewer renders) depends on that and should be confirmed with
  a profile.

---

## Appendix: findings ledger

Severity shown as `reviewer → verifier` where an independent verifier re-checked it.

| Sev | Area | Finding |
|-----|------|---------|
| HIGH→MED | #4 | `watchTickMsg` no generation guard → focus regain spawns parallel tick chain |
| HIGH→MED | #5 | Idle-exit snap-back permanently doubles the watch-tick chain |
| HIGH (confirmed) | #7 | QoS routing inert in all released binaries (`CGO_ENABLED=0`); still pays thread-pinning cost |
| HIGH (confirmed) | bench | Single run, no repetition/variance — can't establish significance |
| HIGH (confirmed) | bench | P/E-core residency `0` in all baselines → PR #7 has no backing data |
| MED | #4 | `WithReportFocus` silent no-op on Terminal.app / ssh / bare tmux |
| MED | #5 | No freshness guard during rollouts / CrashLoopBackOff |
| MED | #7 | Permanently-locked OS thread per informer; unbounded under auto/always |
| MED | #7 | No runtime/config/battery opt-out for E-core routing |
| MED | bench | powermetrics residency zero without sudo |
| MED | bench | `active-scripted` baseline shows anomalous ticker rates |
| MED | merge | E-core QoS no-op in shipped `CGO_ENABLED=0` binary |
| (added) | #4/#5 | idle/blur can *speed up* a configured interval > 30s — use `max(...)` |
| MED | #3 | Final flush skipped on `os.Exit` paths (force-quit / error) |
| LOW | #3 | No periodic flush — crash loses the run |
| LOW | #3 | Dev-only probe linked into release binary |
| LOW | #4 | Focus refresh not gated by overlays/exec/PTY |
| LOW | bench | `-reps` advisory only — silent single run |
| INFO ✅ | #3 | Off-by-default, race-safety, synchronous Stop verified correct |
| INFO ✅ | #3 | `energy.Tick` genuinely wired into ~24 sites; tick metric carries real data |
| INFO ✅ | merge | Fork throttle survived merge; correctly wired with `shuttingDown` guard |
| INFO ✅ | merge | Shutdown drain orders cancel-before-wait; no deadlock vs QoS-locked goroutines |
