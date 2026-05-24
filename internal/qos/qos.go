// Package qos exposes a thin helper for tagging goroutines with a
// macOS QoS (quality-of-service) class so the kernel can route them
// to efficiency cores. On non-darwin platforms, or on darwin built
// without cgo (CGO_ENABLED=0), this package is a no-op stub.
package qos

import "runtime"

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

// RunWith locks the calling goroutine to a dedicated OS thread, sets
// the thread's QoS class to the requested value, runs fn, and returns.
// The thread is intentionally never unlocked: when the goroutine exits,
// the Go runtime destroys the OS thread, which guarantees the QoS class
// cannot leak to an unrelated goroutine via thread reuse.
//
// On non-darwin or no-cgo builds, the QoS step is a no-op and only the
// thread lock + fn execution happen — semantically equivalent but
// energetically neutral.
func RunWith(class QoSClass, fn func()) {
	runtime.LockOSThread()
	setQoS(class)
	fn()
}
