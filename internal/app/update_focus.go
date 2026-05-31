package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// updateBlur handles tea.BlurMsg. It flips the focus flag so the next
// time a watch tick is scheduled, activeWatchInterval returns the
// blurred interval. We deliberately do not try to cancel any in-flight
// watch tick -- Bubble Tea has no tick-cancellation primitive and a
// single stale fire at the old interval is harmless.
func (m Model) updateBlur(_ tea.BlurMsg) (tea.Model, tea.Cmd) {
	m.focused = false
	return m, nil
}

// updateFocus handles tea.FocusMsg. Snap-back UX: immediately refresh
// the current view so the user sees up-to-date data, and reschedule the
// watch tick at the configured (foreground) interval. A stale tick
// scheduled with blurredWatchInterval may still fire later; its
// refreshCurrentLevel call is redundant but cheap.
func (m Model) updateFocus(_ tea.FocusMsg) (tea.Model, tea.Cmd) {
	m.focused = true
	// Regaining focus counts as user activity: reset the idle clock so
	// activeWatchInterval() returns the foreground interval immediately.
	m.lastInputAt = time.Now()
	// suppressBgtasks mirrors updateWatchTick's pattern: trackBgTask
	// captures the flag synchronously at command construction, so we
	// only need it true for the duration of refreshCurrentLevel().
	// Reset before return so it doesn't leak into subsequent Updates.
	m.suppressBgtasks = true
	// Start a fresh chain (nextWatchTick bumps watchTickGen) so any tick
	// still in flight from the blurred period is retired instead of
	// running alongside this one.
	cmd := tea.Batch(m.refreshCurrentLevel(), m.nextWatchTick(m.activeWatchInterval()))
	m.suppressBgtasks = false
	return m, cmd
}
