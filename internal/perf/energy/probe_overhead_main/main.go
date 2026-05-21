// Tiny stand-in process for the probe overhead test. Calls energy.Tick on
// a tight loop so the wrapper's hot path is exercised at high frequency.
package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/janosmiko/lfk/internal/perf/energy"
)

func main() {
	p, _ := energy.Start()
	energy.SetGlobal(p)
	defer func() { _ = p.Stop() }()

	stop := time.After(20 * time.Second)
	for {
		select {
		case <-stop:
			return
		default:
		}
		cmd := energy.Tick("overhead-loop", 5*time.Millisecond, func(t time.Time) tea.Msg { return t })
		_ = cmd()
	}
}
