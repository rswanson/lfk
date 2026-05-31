// Package qos exposes a thin helper for tagging goroutines with a
// macOS QoS (quality-of-service) class so the kernel can route them
// to efficiency cores. On non-darwin platforms, or on darwin built
// without cgo (CGO_ENABLED=0), this package is a no-op: RunWith simply
// runs fn without reserving a dedicated OS thread.
package qos

import (
	"os"
	"runtime"
)

// QoSClass mirrors macOS's qos_class_t (defined in <sys/qos.h>).
// The constants below are the only values the lfk codebase should
// pass; adding USER_INTERACTIVE / USER_INITIATED / DEFAULT here would
// only invite misuse.
type QoSClass int

const (
	// Utility is appropriate for long-running but visible work whose
	// latency the user might eventually notice. Informer watch loops
	// are the canonical case: data freshness depends on them, but the
	// user is not interactively waiting on each event.
	Utility QoSClass = 17

	// Background is for truly invisible work whose precise timing the
	// user does not care about. The energy probe sampler is the
	// canonical case.
	Background QoSClass = 9
)

// disableEnv opts out of QoS tagging at runtime when set to "off". It
// only matters on a build where QoS is actually applied (darwin+cgo);
// on every other build RunWith is already a no-op. Provided as an
// escape hatch in case routing the informer to E-cores makes the
// resource view feel laggy under CPU pressure or on battery.
const disableEnv = "LFK_QOS"

// RunWith runs fn, optionally on a goroutine tagged with a macOS QoS
// class. When QoS can actually be applied — a darwin+cgo build with
// LFK_QOS unset — it locks the calling goroutine to a dedicated OS
// thread and sets that thread's QoS class. The thread is intentionally
// never unlocked: when the goroutine exits the runtime destroys the OS
// thread, which guarantees the QoS class cannot leak to an unrelated
// goroutine via thread reuse.
//
// On any other build (non-darwin, no-cgo, or LFK_QOS=off) the QoS step
// is a no-op, so RunWith does NOT lock a thread — locking one would
// permanently reserve an OS thread for zero scheduling benefit (the
// bug this guard fixes: release binaries are CGO_ENABLED=0, so the old
// unconditional LockOSThread stranded one OS thread per informer with
// no QoS payoff). qosSupported is a per-build constant, so the lock
// path is compiled out entirely where it cannot help.
func RunWith(class QoSClass, fn func()) {
	if qosSupported && os.Getenv(disableEnv) != "off" {
		runtime.LockOSThread()
		setQoS(class)
	}
	fn()
}
