package app

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdate_BlurMsgFlipsFocused(t *testing.T) {
	m := Model{focused: true, watchInterval: 2 * time.Second}
	out, cmd := m.Update(tea.BlurMsg{})
	mod, ok := out.(Model)
	if !ok {
		t.Fatalf("Update returned %T, want Model", out)
	}
	if mod.focused {
		t.Errorf("focused after BlurMsg = true, want false")
	}
	if cmd != nil {
		t.Errorf("BlurMsg cmd = %v, want nil (no immediate side-effect)", cmd)
	}
}

func TestUpdate_FocusMsgFlipsFocusedAndEmitsRefresh(t *testing.T) {
	m := Model{focused: false, watchInterval: 2 * time.Second, watchMode: true}
	out, cmd := m.Update(tea.FocusMsg{})
	mod, ok := out.(Model)
	if !ok {
		t.Fatalf("Update returned %T, want Model", out)
	}
	if !mod.focused {
		t.Errorf("focused after FocusMsg = false, want true")
	}
	if cmd == nil {
		t.Fatal("FocusMsg cmd = nil; expected tea.Batch(refresh, scheduleWatchTick)")
	}
	// We don't assert on cmd's internals; behaviour is verified by the
	// fact that a non-nil command is returned and refreshCurrentLevel is
	// covered by its own tests.
}

func TestUpdate_FocusMsgStillEmitsCmdWhenWatchModeOff(t *testing.T) {
	// Even without watch mode, FocusMsg should still trigger a refresh
	// so the visible data is current when the user returns.
	m := Model{focused: false, watchInterval: 2 * time.Second, watchMode: false}
	_, cmd := m.Update(tea.FocusMsg{})
	if cmd == nil {
		t.Error("FocusMsg with watchMode=false: cmd should still trigger refresh")
	}
}

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
