package energy

import (
	"sync/atomic"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// globalProbe is the process-wide probe set by SetGlobal so Tick can be a
// top-level function callable from anywhere. atomic.Pointer keeps the hot
// path lock-free.
var globalProbe atomic.Pointer[Probe]

// SetGlobal registers p as the process-wide probe used by Tick.
// Pass nil to clear it (used in tests).
func SetGlobal(p *Probe) { globalProbe.Store(p) }

// Tick is a labelled, instrumented replacement for tea.Tick. When the
// probe is disabled (or unset), it is a thin pass-through. When enabled,
// it increments the tick counter for the given label.
func Tick(label string, d time.Duration, fn func(time.Time) tea.Msg) tea.Cmd {
	return tea.Tick(d, func(now time.Time) tea.Msg {
		if p := globalProbe.Load(); p != nil && p.enabled {
			p.tickMu.Lock()
			p.tickCounts[label]++
			p.tickMu.Unlock()
		}
		return fn(now)
	})
}
