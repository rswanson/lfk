package main

import (
	"bufio"
	"context"
	"io"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

// topStreamSampler runs a `top` (or compatible) process in continuous mode
// and accumulates per-block samples filtered to a single PID. The first
// block is discarded because top's initial IDLEW value is cumulative since
// boot rather than a delta.
type topStreamSampler struct {
	bin      string
	args     []string
	interval time.Duration

	mu        sync.Mutex
	started   bool
	stopped   bool
	cmd       *exec.Cmd
	stdout    io.ReadCloser
	collected []topSample
	wg        sync.WaitGroup
}

// newTopStreamSampler builds a sampler that will run `bin args...` on
// Start. For real use, pass topBin() and topArgs(); tests can pass a fake.
func newTopStreamSampler(bin string, args []string, interval time.Duration) *topStreamSampler {
	return &topStreamSampler{bin: bin, args: args, interval: interval}
}

// topBin returns the `top` binary to invoke on the current OS.
// Wired into main.go in Task 3; not yet called within this file.
//
//nolint:unused // called by main.go (Task 3 wires the streaming sampler into run())
func topBin() string { return "top" }

// topArgs returns the macOS top args for continuous, 1s-interval sampling
// with the columns the parser expects.
// Wired into main.go in Task 3; not yet called within this file.
//
//nolint:unused // called by main.go (Task 3 wires the streaming sampler into run())
func topArgs(interval time.Duration) []string {
	// -l 0: log forever; -s N: sample every N seconds; -stats: fixed columns.
	secs := max(int(interval.Seconds()), 1)
	return []string{"-l", "0", "-s", strconv.Itoa(secs), "-stats", "pid,cpu,power,idlew"}
}

// Start launches the subprocess and begins collecting samples for the given PID.
// Returns an error if the process cannot start.
func (s *topStreamSampler) Start(ctx context.Context, pid int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.started {
		return nil
	}
	s.started = true

	cmd := exec.CommandContext(ctx, s.bin, s.args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	s.cmd = cmd
	s.stdout = stdout

	s.wg.Add(1)
	go s.collect(pid)
	return nil
}

// collect reads stdout one block at a time, parsing each as a top dump
// and appending PID-filtered rows. Discards the very first block.
func (s *topStreamSampler) collect(pid int) {
	defer s.wg.Done()

	br := bufio.NewReader(s.stdout)
	var (
		block    []byte
		blockIdx int
		line     []byte
		err      error
	)
	for {
		line, err = br.ReadBytes('\n')
		if len(line) > 0 {
			// A new block starts whenever we hit a "Processes:" header.
			// Treat that as a block boundary: flush the previous block.
			if isBlockStart(line) && len(block) > 0 {
				s.handleBlock(block, blockIdx, pid)
				blockIdx++
				block = block[:0]
			}
			block = append(block, line...)
		}
		if err != nil {
			break
		}
	}
	if len(block) > 0 {
		s.handleBlock(block, blockIdx, pid)
	}
}

func isBlockStart(line []byte) bool {
	const prefix = "Processes:"
	if len(line) < len(prefix) {
		return false
	}
	for i := range len(prefix) {
		if line[i] != prefix[i] {
			return false
		}
	}
	return true
}

func (s *topStreamSampler) handleBlock(block []byte, idx int, pid int) {
	if idx == 0 {
		return // warmup block: discard
	}
	samples, err := parseTopOutputForPID(string(block), pid)
	if err != nil || len(samples) == 0 {
		return
	}
	s.mu.Lock()
	s.collected = append(s.collected, samples...)
	s.mu.Unlock()
}

// Stop kills the subprocess and returns the collected samples. Safe to
// call more than once.
func (s *topStreamSampler) Stop() []topSample {
	s.mu.Lock()
	if s.stopped || !s.started {
		out := append([]topSample(nil), s.collected...)
		s.stopped = true
		s.mu.Unlock()
		return out
	}
	s.stopped = true
	cmd := s.cmd
	s.mu.Unlock()

	if cmd != nil && cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
	s.wg.Wait()
	_ = cmd.Wait()

	s.mu.Lock()
	out := append([]topSample(nil), s.collected...)
	s.mu.Unlock()
	return out
}
