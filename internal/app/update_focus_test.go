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
	m := Model{watchInterval: 3 * time.Second, focused: true, lastInputAt: time.Now()}
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

func TestSuppressBgtasksFlagDoesNotLeakAfterFocusMsg(t *testing.T) {
	// Mirror of TestSuppressBgtasksFlagDoesNotLeakAfterWatchTick — guards
	// the same value-receiver pattern in updateFocus.
	m := Model{focused: false, watchInterval: 2 * time.Second, watchMode: true}
	if m.suppressBgtasks {
		t.Fatal("precondition: suppressBgtasks should start false")
	}

	out, _ := m.Update(tea.FocusMsg{})
	updated, ok := out.(Model)
	if !ok {
		t.Fatalf("Update returned %T, want Model", out)
	}
	if updated.suppressBgtasks {
		t.Error("suppressBgtasks leaked: still true on the model returned from FocusMsg")
	}
}

func TestModel_ForegroundIdle_FalseRightAfterConstruction(t *testing.T) {
	m := Model{lastInputAt: time.Now()}
	if m.foregroundIdle() {
		t.Error("foregroundIdle returned true immediately after construction")
	}
}

func TestModel_ForegroundIdle_TrueAfterThreshold(t *testing.T) {
	m := Model{lastInputAt: time.Now().Add(-2 * foregroundIdleThreshold)}
	if !m.foregroundIdle() {
		t.Errorf("foregroundIdle returned false after %v of inactivity", 2*foregroundIdleThreshold)
	}
}

func TestModel_ForegroundIdleThresholdIs120s(t *testing.T) {
	if foregroundIdleThreshold != 120*time.Second {
		t.Errorf("foregroundIdleThreshold = %v, want 120s (spec)", foregroundIdleThreshold)
	}
}

func TestModel_ActiveWatchInterval_IdleAndFocusedReturnsBlurred(t *testing.T) {
	m := Model{
		watchInterval: 2 * time.Second,
		focused:       true,
		lastInputAt:   time.Now().Add(-2 * foregroundIdleThreshold),
	}
	if got := m.activeWatchInterval(); got != blurredWatchInterval {
		t.Errorf("idle+focused: got %v, want %v", got, blurredWatchInterval)
	}
}

func TestModel_ActiveWatchInterval_FocusedAndRecentReturnsConfigured(t *testing.T) {
	m := Model{
		watchInterval: 2 * time.Second,
		focused:       true,
		lastInputAt:   time.Now(),
	}
	if got := m.activeWatchInterval(); got != m.watchInterval {
		t.Errorf("focused+recent: got %v, want %v", got, m.watchInterval)
	}
}

func TestModel_SnapBackIfIdle_NilWhenNotIdle(t *testing.T) {
	m := Model{focused: true, lastInputAt: time.Now()}
	if cmd := m.snapBackIfIdle(); cmd != nil {
		t.Errorf("snapBackIfIdle = %v, want nil when not idle", cmd)
	}
}

func TestModel_SnapBackIfIdle_NilWhenBlurred(t *testing.T) {
	// PR-4 owns the blurred path; this helper should defer to it.
	m := Model{
		focused:     false,
		lastInputAt: time.Now().Add(-2 * foregroundIdleThreshold),
	}
	if cmd := m.snapBackIfIdle(); cmd != nil {
		t.Errorf("snapBackIfIdle = %v, want nil when blurred", cmd)
	}
}

func TestModel_SnapBackIfIdle_NonNilWhenFocusedAndIdle(t *testing.T) {
	m := Model{
		focused:     true,
		lastInputAt: time.Now().Add(-2 * foregroundIdleThreshold),
	}
	if cmd := m.snapBackIfIdle(); cmd == nil {
		t.Error("snapBackIfIdle = nil, want non-nil when foreground-idle")
	}
}

func TestUpdate_KeyMsgUpdatesLastInputAt(t *testing.T) {
	stale := time.Now().Add(-2 * foregroundIdleThreshold)
	m := Model{focused: true, watchInterval: 2 * time.Second, lastInputAt: stale}
	out, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	mod, ok := out.(Model)
	if !ok {
		t.Fatalf("Update returned %T, want Model", out)
	}
	if !mod.lastInputAt.After(stale) {
		t.Errorf("lastInputAt = %v, want > %v", mod.lastInputAt, stale)
	}
}

func TestUpdate_KeyMsgAfterIdleEmitsSnapBackCmd(t *testing.T) {
	m := Model{
		focused:       true,
		watchMode:     true,
		watchInterval: 2 * time.Second,
		lastInputAt:   time.Now().Add(-2 * foregroundIdleThreshold),
	}
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	if cmd == nil {
		t.Fatal("KeyMsg after idle: cmd was nil; expected snap-back batched with handler cmd")
	}
	// We don't introspect cmd internals -- coverage is via integration
	// with refreshCurrentLevel + scheduleWatchTick which have their own tests.
}

func TestUpdate_FocusMsgUpdatesLastInputAt(t *testing.T) {
	// Regression guard: FocusMsg must set lastInputAt so a user who
	// returns to a window they left idle for >120s isn't immediately
	// re-marked as idle.
	stale := time.Now().Add(-2 * foregroundIdleThreshold)
	m := Model{focused: false, watchInterval: 2 * time.Second, lastInputAt: stale}
	out, _ := m.Update(tea.FocusMsg{})
	mod, ok := out.(Model)
	if !ok {
		t.Fatalf("Update returned %T, want Model", out)
	}
	if !mod.lastInputAt.After(stale) {
		t.Errorf("lastInputAt after FocusMsg = %v, want > %v", mod.lastInputAt, stale)
	}
	if mod.foregroundIdle() {
		t.Error("foregroundIdle = true right after FocusMsg; should be false")
	}
}
