//go:build darwin && cgo

package qos

import (
	"testing"
)

func TestRunWith_AppliesUtilityClass(t *testing.T) {
	got := make(chan QoSClass, 1)
	go RunWith(Utility, func() {
		got <- currentQoS()
	})
	if v := <-got; v != Utility {
		t.Errorf("currentQoS inside RunWith(Utility) = %d, want %d", v, Utility)
	}
}

func TestRunWith_AppliesBackgroundClass(t *testing.T) {
	got := make(chan QoSClass, 1)
	go RunWith(Background, func() {
		got <- currentQoS()
	})
	if v := <-got; v != Background {
		t.Errorf("currentQoS inside RunWith(Background) = %d, want %d", v, Background)
	}
}
