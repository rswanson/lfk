// Command energy-bench drives lfk against a fake apiserver and produces an
// energy/perf report. See docs/perf/energy-bench.md for usage.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	tmp, err := os.MkdirTemp("", "energy-bench-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	kc := tmp + "/kubeconfig"
	if err := writeKubeconfig(kc, "http://127.0.0.1:0"); err != nil {
		return err
	}
	fmt.Printf("energy-bench: scaffold ok, scenario=%s fixture=%s reps=%d\n",
		o.scenario, o.fixture, o.repetitions)
	return nil
}
