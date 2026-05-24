//go:build darwin && cgo

package qos

/*
#include <pthread.h>
#include <sys/qos.h>
*/
import "C"

func setQoS(class QoSClass) {
	// The second arg is "relative priority within class"; 0 is the
	// only value Apple's docs really define. Higher relative priorities
	// pull the thread back toward the next class up, defeating the
	// purpose.
	C.pthread_set_qos_class_self_np(C.qos_class_t(class), 0)
}

// currentQoS reads the current thread's QoS class. Test-only helper
// exposed for qos_darwin_test.go; not part of the public surface.
func currentQoS() QoSClass {
	var c C.qos_class_t
	C.pthread_get_qos_class_np(C.pthread_self(), &c, nil)
	return QoSClass(c)
}
