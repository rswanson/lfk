// Command energy-bench drives lfk against a fake apiserver and produces an
// energy/perf report. See docs/perf/energy-bench.md for usage.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/janosmiko/lfk/internal/perf/energy/fakeapi"
)

type opts struct {
	binary      string
	scenario    string
	fixture     string
	contexts    int
	repetitions int
	durationSec int
	outDir      string
	useSudo     bool
	updateBase  bool
}

func main() {
	var o opts
	flag.StringVar(&o.binary, "binary", "./lfk", "path to lfk binary")
	flag.StringVar(&o.scenario, "scenario", "idle-foreground",
		"scenario: idle-foreground | idle-background | active-scripted")
	flag.StringVar(&o.fixture, "fixture", "medium", "fixture: small | medium | large")
	flag.IntVar(&o.contexts, "contexts", 3, "number of kubeconfig contexts")
	flag.IntVar(&o.repetitions, "reps", 3, "number of repetitions per scenario")
	flag.IntVar(&o.durationSec, "duration", 90, "scenario duration in seconds")
	flag.StringVar(&o.outDir, "out", "testdata/energy-reports", "output directory")
	flag.BoolVar(&o.useSudo, "powermetrics", false, "use powermetrics (requires sudo)")
	flag.BoolVar(&o.updateBase, "update-baseline", false, "write current run as new baseline")
	flag.Parse()

	if err := run(o); err != nil {
		log.Fatalf("energy-bench: %v", err)
	}
}

func run(o opts) error {
	if _, err := os.Stat(o.binary); err != nil {
		return fmt.Errorf("binary %q not found: %w", o.binary, err)
	}

	fx := selectFixture(o.fixture)
	srv := fakeapi.New(fx).Start()
	defer srv.Close()

	tmp, err := os.MkdirTemp("", "energy-bench-")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }()
	kcPath := filepath.Join(tmp, "kubeconfig")
	if err := writeKubeconfig(kcPath, srv.URL); err != nil {
		return err
	}

	dataDir := filepath.Join(tmp, "data")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}
	env := append(os.Environ(),
		"LFK_ENERGY_PROBE=1",
		"LFK_DATA_DIR="+dataDir,
		"KUBECONFIG="+kcPath,
	)

	ctx := context.Background()
	sampler := newTopStreamSampler(topBin(), topArgs(time.Second), time.Second)
	defer func() { _ = sampler.Stop() }() // ensure top is killed even if a scenario errors

	cfg := scenarioConfig{
		binary:      o.binary,
		args:        nil,
		env:         env,
		durationSec: o.durationSec,
		onProcessStart: func(pid int) {
			if err := sampler.Start(ctx, pid); err != nil {
				log.Printf("warning: top sampler failed to start: %v (wakeups/energy will be zero)", err)
			}
		},
	}

	switch o.scenario {
	case "idle-foreground":
		err = runIdleForeground(ctx, cfg)
	case "idle-background":
		err = runIdleBackground(ctx, cfg, "Ghostty")
	case "active-scripted":
		err = runActiveScripted(ctx, cfg)
	default:
		return fmt.Errorf("unknown scenario %q", o.scenario)
	}
	if err != nil {
		return err
	}

	// Stop is idempotent; an explicit call here returns the collected
	// samples while the deferred Stop above guarantees cleanup on the
	// error paths above.
	topSamples := sampler.Stop()
	if len(topSamples) == 0 {
		log.Printf("warning: no PID-filtered top samples collected; wakeups/energy will be zero")
	}

	var pm powermetricsSample
	if o.useSudo {
		pmOut, pmErr := exec.Command("sudo", "powermetrics",
			"--samplers", "cpu_power,tasks", "-i", "1000", "-n", "1").Output()
		if pmErr != nil {
			log.Printf("warning: powermetrics failed: %v (P/E-core residency will be zero)", pmErr)
		}
		var perr error
		pm, perr = parsePowermetricsOutput(string(pmOut))
		if perr != nil {
			log.Printf("warning: parsePowermetricsOutput failed: %v", perr)
		}
	}

	// Stop the probe (it is inside the lfk subprocess and was killed when
	// the PTY closed); read whatever JSONL it wrote to dataDir/energy/.
	probeAgg, perr := ingestProbe(filepath.Join(dataDir, "energy"))
	if perr != nil {
		log.Printf("warning: probe JSONL ingest failed: %v", perr)
	}

	r := buildReport(reportInputs{
		Scenario:    o.scenario,
		Fixture:     o.fixture,
		Contexts:    o.contexts,
		DurationSec: o.durationSec,
		Reps:        o.repetitions,
		Top:         topSamples,
		Power:       pm,
		Probe:       probeAgg,
	})

	baseDir := "testdata/baselines"
	basePath := filepath.Join(baseDir, o.scenario+".json")
	baseline, _ := readBaseline(basePath)

	stamp := time.Now().Format("20060102-150405")
	runDir := filepath.Join(o.outDir, fmt.Sprintf("%s-%s", o.scenario, stamp))
	if err := os.MkdirAll(runDir, 0o755); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(r, "", "  "); err == nil {
		_ = os.WriteFile(filepath.Join(runDir, "report.json"), data, 0o644)
	}
	_ = os.WriteFile(filepath.Join(runDir, "report.md"), []byte(renderMarkdown(r, baseline)), 0o644)

	if o.updateBase {
		if err := os.MkdirAll(baseDir, 0o755); err != nil {
			return err
		}
		if err := writeBaseline(basePath, r); err != nil {
			return err
		}
	}
	fmt.Printf("energy-bench: report at %s\n", runDir)
	return nil
}

func selectFixture(name string) fakeapi.Fixture {
	switch name {
	case "small":
		return fakeapi.SmallFixture()
	case "large":
		return fakeapi.LargeFixture()
	default:
		return fakeapi.MediumFixture()
	}
}
