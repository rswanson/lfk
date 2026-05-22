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
	if got := mod.activeWatchInterval(); got != m.watchInterval {
		t.Errorf("activeWatchInterval after FocusMsg = %v, want %v (foreground interval)", got, m.watchInterval)
	}
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

func TestUpdate_BlurThenFocusReturnsToForegroundInterval(t *testing.T) {
	// Sequence flow: focused start, BlurMsg flips to blurred interval,
	// FocusMsg snaps back to the configured foreground interval.
	m := Model{focused: true, watchInterval: 2 * time.Second, watchMode: true}

	afterBlur, _ := m.Update(tea.BlurMsg{})
	mBlur, ok := afterBlur.(Model)
	if !ok {
		t.Fatalf("after BlurMsg: got %T, want Model", afterBlur)
	}
	if got := mBlur.activeWatchInterval(); got != blurredWatchInterval {
		t.Errorf("activeWatchInterval after BlurMsg = %v, want %v", got, blurredWatchInterval)
	}

	afterFocus, _ := mBlur.Update(tea.FocusMsg{})
	mFocus, ok := afterFocus.(Model)
	if !ok {
		t.Fatalf("after FocusMsg: got %T, want Model", afterFocus)
	}
	if got := mFocus.activeWatchInterval(); got != m.watchInterval {
		t.Errorf("activeWatchInterval after FocusMsg = %v, want %v", got, m.watchInterval)
	}
}
