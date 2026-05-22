package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// updateBlur handles tea.BlurMsg. It flips the focus flag so the next
// time a watch tick is scheduled, activeWatchInterval returns
// blurredWatchInterval. We deliberately do not try to cancel any
// in-flight watch tick -- Bubble Tea has no tick-cancellation primitive
// and a single stale fire at the old interval is harmless.
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
	return m, tea.Batch(m.refreshCurrentLevel(), scheduleWatchTick(m.activeWatchInterval()))
}
