package main

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/creack/pty"
)

type ptySession struct {
	cmd *exec.Cmd
	tty io.ReadWriteCloser

	mu  sync.Mutex
	buf strings.Builder
}

func startPTY(ctx context.Context, bin string, args, env []string) (*ptySession, error) {
	c := exec.CommandContext(ctx, bin, args...)
	c.Env = env
	f, err := pty.Start(c)
	if err != nil {
		return nil, err
	}
	s := &ptySession{cmd: c, tty: f}
	go s.drain()
	return s, nil
}

func (s *ptySession) drain() {
	buf := make([]byte, 4096)
	for {
		n, err := s.tty.Read(buf)
		if n > 0 {
			s.mu.Lock()
			s.buf.Write(buf[:n])
			s.mu.Unlock()
		}
		if err != nil {
			return
		}
	}
}

func (s *ptySession) SendString(in string) error {
	_, err := s.tty.Write([]byte(in))
	return err
}

func (s *ptySession) ReadUntil(needle string, timeout time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		s.mu.Lock()
		got := s.buf.String()
		s.mu.Unlock()
		if strings.Contains(got, needle) {
			return got, nil
		}
		time.Sleep(10 * time.Millisecond)
	}
	return "", errors.New("timeout waiting for " + needle)
}

// probeFlushGrace is the time we give lfk's SIGUSR1 handler to write the
// probe JSONL file before we SIGKILL. The handler does one file write, so
// this is conservative; raise it if probe data turns up missing.
const probeFlushGrace = 250 * time.Millisecond

func (s *ptySession) Close() {
	if s.cmd.Process != nil {
		// SIGUSR1 triggers lfk's probe.FlushOnSignal handler, which
		// writes the probe ring buffer to JSONL. SIGKILL skips defers,
		// so without this the probe never flushes when driven via PTY.
		_ = s.cmd.Process.Signal(syscall.SIGUSR1)
		time.Sleep(probeFlushGrace)
		_ = s.cmd.Process.Kill()
	}
	_ = s.tty.Close()
	_ = s.cmd.Wait()
}

// PID returns the subprocess PID, or 0 if the process has not started.
// The PID may refer to an already-exited process if the subprocess exits
// before the caller uses the value; callers must handle that case.
func (s *ptySession) PID() int {
	if s.cmd == nil || s.cmd.Process == nil {
		return 0
	}
	return s.cmd.Process.Pid
}
