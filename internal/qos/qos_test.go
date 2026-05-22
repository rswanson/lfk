package qos

import (
	"testing"
)

func TestRunWith_CallsFnExactlyOnce(t *testing.T) {
	calls := 0
	done := make(chan struct{})
	go func() {
		defer close(done)
		RunWith(Utility, func() {
			calls++
		})
	}()
	<-done
	if calls != 1 {
		t.Errorf("RunWith called fn %d times, want 1", calls)
	}
}

func TestRunWith_BackgroundClassValuesAreStable(t *testing.T) {
	// Values match macOS qos_class_t. Locking these in case anyone
	// touches the constants by accident.
	if Utility != 17 {
		t.Errorf("Utility = %d, want 17 (QOS_CLASS_UTILITY)", Utility)
	}
	if Background != 9 {
		t.Errorf("Background = %d, want 9 (QOS_CLASS_BACKGROUND)", Background)
	}
}
