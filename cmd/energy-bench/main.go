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

	// 1. Start fake apiserver.
	fx := selectFixture(o.fixture)
	srv := fakeapi.New(fx).Start()
	defer srv.Close()

	// 2. Throwaway kubeconfig.
	tmp, err := os.MkdirTemp("", "energy-bench-")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }()
	kcPath := filepath.Join(tmp, "kubeconfig")
	if err := writeKubeconfig(kcPath, srv.URL); err != nil {
		return err
	}

	// 3. Environment for the target process.
	dataDir := filepath.Join(tmp, "data")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}
	env := append(os.Environ(),
		"LFK_ENERGY_PROBE=1",
		"LFK_DATA_DIR="+dataDir,
		"KUBECONFIG="+kcPath,
	)

	// 4. Run the scenario.
	cfg := scenarioConfig{
		binary:      o.binary,
		args:        nil,
		env:         env,
		durationSec: o.durationSec,
	}
	ctx := context.Background()
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

	// 5. Sample top/powermetrics. For Phase 1 we take a short tail
	//    sample. Future revisions can run sampling concurrently with
	//    the scenario.
	topOut, _ := exec.Command("top", "-l", "1", "-stats", "pid,cpu,power,idlew").Output()
	topSamples, _ := parseTopOutput(string(topOut))
	var pm powermetricsSample
	if o.useSudo {
		pmOut, _ := exec.Command("sudo", "powermetrics", "--samplers", "cpu_power,tasks", "-i", "1000", "-n", "1").Output()
		pm, _ = parsePowermetricsOutput(string(pmOut))
	}

	// 6. Build report.
	r := buildReport(reportInputs{
		Scenario:    o.scenario,
		Fixture:     o.fixture,
		Contexts:    o.contexts,
		DurationSec: o.durationSec,
		Reps:        o.repetitions,
		Top:         topSamples,
		Power:       pm,
	})

	// 7. Compare to baseline.
	baseDir := "testdata/baselines"
	basePath := filepath.Join(baseDir, o.scenario+".json")
	baseline, _ := readBaseline(basePath)

	// 8. Write outputs.
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
