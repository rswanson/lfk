// Tiny stand-in process for the probe overhead test.
//
// This is an ADVERSARIAL stress test: it calls energy.Tick and immediately
// executes the returned tea.Cmd in a tight loop, so the wrapper's hot path
// is hit hundreds of times per second. Real lfk schedules tea.Cmds through
// the bubbletea runtime at much lower frequencies (5ms-5s intervals,
// scheduler-blocked between fires), so this measures a worst-case
// upper-bound — not steady-state production overhead.
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
