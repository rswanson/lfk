package main

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"strings"
	"sync"
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

func (s *ptySession) Close() {
	_ = s.tty.Close()
	if s.cmd.Process != nil {
		_ = s.cmd.Process.Kill()
	}
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
