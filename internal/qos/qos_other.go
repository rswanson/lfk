//go:build !(darwin && cgo)

package qos

// qosSupported is false on builds without the cgo QoS syscall, so RunWith
// skips LockOSThread entirely (no dedicated OS thread is reserved).
const qosSupported = false

func setQoS(_ QoSClass) {} // no-op on non-darwin or no-cgo builds
