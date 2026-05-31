package app

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// These tests pin the generation-token contract that keeps exactly one
// watch-tick chain alive. Before the fix, watchTickMsg carried no
// generation and tea.Tick cannot be cancelled, so every focus regain /
// idle snap-back started a NEW self-re-arming chain on top of the old
// one, multiplying refreshes under the very churn the throttle targets.

func TestUpdateWatchTick_DropsStaleGeneration(t *testing.T) {
	m := Model{
		watchMode:     true,
		watchTickGen:  5,
		watchInterval: 2 * time.Second,
		focused:       true,
		lastInputAt:   time.Now(),
	}
	_, cmd := m.updateWatchTick(watchTickMsg{gen: 4})
	if cmd != nil {
		t.Error("stale-generation watch tick returned a cmd; want nil so the orphaned chain dies")
	}
}

func TestUpdateWatchTick_HonorsCurrentGeneration(t *testing.T) {
	m := Model{
		watchMode:     true,
		watchTickGen:  5,
		watchInterval: 2 * time.Second,
		focused:       true,
		lastInputAt:   time.Now(),
	}
	_, cmd := m.updateWatchTick(watchTickMsg{gen: 5})
	if cmd == nil {
		t.Error("current-generation watch tick returned nil; want a cmd so the one live chain continues")
	}
}

func TestUpdateFocus_AdvancesWatchTickGeneration(t *testing.T) {
	m := Model{focused: false, watchMode: true, watchInterval: 2 * time.Second, watchTickGen: 7}
	out, _ := m.Update(tea.FocusMsg{})
	mod := out.(Model)
	if mod.watchTickGen <= 7 {
		t.Errorf("watchTickGen after FocusMsg = %d, want > 7 (focus starts a fresh chain)", mod.watchTickGen)
	}
}

func TestUpdateBlur_DoesNotAdvanceWatchTickGeneration(t *testing.T) {
	// Blur must NOT start a new chain: the single live chain simply slows
	// to blurredWatchInterval on its next re-arm (activeWatchInterval reads
	// m.focused). Bumping here would either orphan the loop or duplicate it.
	m := Model{focused: true, watchMode: true, watchInterval: 2 * time.Second, watchTickGen: 7}
	out, _ := m.Update(tea.BlurMsg{})
	mod := out.(Model)
	if mod.watchTickGen != 7 {
		t.Errorf("watchTickGen after BlurMsg = %d, want 7 unchanged", mod.watchTickGen)
	}
}

func TestUpdate_FocusFlapLeavesSingleLiveWatchChain(t *testing.T) {
	// Headline regression: drive blur->focus->blur->focus and prove the
	// first focus regain's chain is retired by the second, so chains do not
	// accrete. The generation captured after the first focus must be stale
	// (dropped) once the second focus has advanced the generation.
	m := Model{focused: true, watchMode: true, watchInterval: 2 * time.Second, lastInputAt: time.Now()}

	var cur tea.Model = m
	gens := make([]uint64, 0, 2)
	for range 2 {
		afterBlur, _ := cur.(Model).Update(tea.BlurMsg{})
		afterFocus, _ := afterBlur.(Model).Update(tea.FocusMsg{})
		gens = append(gens, afterFocus.(Model).watchTickGen)
		cur = afterFocus
	}
	final := cur.(Model)

	if gens[0] == gens[1] {
		t.Fatalf("watchTickGen did not advance across focus regains: %d == %d", gens[0], gens[1])
	}

	// The first focus regain's chain is now stale -> its tick must be dropped.
	if _, cmd := final.updateWatchTick(watchTickMsg{gen: gens[0]}); cmd != nil {
		t.Errorf("first-regain tick (gen %d) honored; want dropped (current gen %d) — chains accreted", gens[0], final.watchTickGen)
	}
	// The current chain's tick is honored -> exactly one chain lives on.
	if _, cmd := final.updateWatchTick(watchTickMsg{gen: final.watchTickGen}); cmd == nil {
		t.Errorf("current tick (gen %d) dropped; the single live chain must continue", final.watchTickGen)
	}
}

func TestUpdate_KeyMsgSnapBackRetiresIdleWatchChain(t *testing.T) {
	// After foreground-idle, a keypress snaps back to the fast cadence by
	// starting a fresh chain; the slow idle chain must be retired so the
	// two do not run in parallel.
	m := Model{
		focused:       true,
		watchMode:     true,
		watchInterval: 2 * time.Second,
		lastInputAt:   time.Now().Add(-2 * foregroundIdleThreshold),
	}
	genBefore := m.watchTickGen

	out, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	after := out.(Model)
	if after.watchTickGen == genBefore {
		t.Fatalf("watchTickGen unchanged (%d) after idle snap-back; the slow chain was not retired", genBefore)
	}
	if _, cmd := after.updateWatchTick(watchTickMsg{gen: genBefore}); cmd != nil {
		t.Error("pre-snap-back idle tick honored; want dropped so the slow chain dies")
	}
}

func TestUpdate_NonIdleKeyMsgDoesNotAdvanceGeneration(t *testing.T) {
	// A keypress while NOT idle must not bump the generation — otherwise
	// every keystroke would retire and rebuild the watch chain.
	m := Model{focused: true, watchMode: true, watchInterval: 2 * time.Second, lastInputAt: time.Now(), watchTickGen: 3}
	out, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	if got := out.(Model).watchTickGen; got != 3 {
		t.Errorf("watchTickGen after non-idle keypress = %d, want 3 unchanged", got)
	}
}
