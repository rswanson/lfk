//go:build !(darwin && cgo)

package qos

func setQoS(_ QoSClass) {} // no-op on non-darwin or no-cgo builds
