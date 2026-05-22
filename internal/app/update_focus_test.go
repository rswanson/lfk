package app

import (
	"testing"
	"time"
)

func TestModel_ActiveWatchInterval_FocusedReturnsConfigured(t *testing.T) {
	m := Model{watchInterval: 3 * time.Second, focused: true}
	if got := m.activeWatchInterval(); got != 3*time.Second {
		t.Errorf("focused: got %v, want 3s", got)
	}
}

func TestModel_ActiveWatchInterval_BlurredReturnsBlurredConst(t *testing.T) {
	m := Model{watchInterval: 3 * time.Second, focused: false}
	if got := m.activeWatchInterval(); got != blurredWatchInterval {
		t.Errorf("blurred: got %v, want %v", got, blurredWatchInterval)
	}
}

func TestModel_ActiveWatchInterval_BlurredConstIs30s(t *testing.T) {
	if blurredWatchInterval != 30*time.Second {
		t.Errorf("blurredWatchInterval = %v, want 30s (spec)", blurredWatchInterval)
	}
}
