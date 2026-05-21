package energy

import (
	"runtime"
	"runtime/metrics"
	"time"
)

// Sample holds a point-in-time snapshot of runtime metrics.
type Sample struct {
	WallNanos       int64             `json:"wall_nanos"`
	Goroutines      int               `json:"goroutines"`
	HeapAllocBytes  uint64            `json:"heap_alloc_bytes"`
	NumGC           uint32            `json:"num_gc"`
	GCPauseTotalNs  uint64            `json:"gc_pause_total_ns"`
	GCCPUSeconds    float64           `json:"gc_cpu_seconds"`
	SchedLatencyP99 float64           `json:"sched_latency_p99_seconds"`
	TickCounts      map[string]uint64 `json:"tick_counts"`
}

const defaultRingCapacity = 4096

func (p *Probe) startSampler() {
	p.ring = make([]Sample, 0, p.ringCap)
	p.wg.Add(1)
	go p.sampleLoop()
}

func (p *Probe) sampleLoop() {
	defer p.wg.Done()
	t := time.NewTicker(p.tickInterval)
	defer t.Stop()
	for {
		select {
		case <-p.stopCh:
			return
		case <-t.C:
			p.appendSample(p.collectSample())
		}
	}
}

func (p *Probe) collectSample() Sample {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	metrics.Read(p.schedSamples)
	metrics.Read(p.gcCPUSamples)

	p99 := 0.0
	if p.schedSamples[0].Value.Kind() == metrics.KindFloat64Histogram {
		if h := p.schedSamples[0].Value.Float64Histogram(); h != nil {
			p99 = histP99(h)
		}
	}
	gcCPU := 0.0
	if p.gcCPUSamples[0].Value.Kind() == metrics.KindFloat64 {
		gcCPU = p.gcCPUSamples[0].Value.Float64()
	}

	return Sample{
		WallNanos:       time.Now().UnixNano(),
		Goroutines:      runtime.NumGoroutine(),
		HeapAllocBytes:  m.HeapAlloc,
		NumGC:           m.NumGC,
		GCPauseTotalNs:  m.PauseTotalNs,
		GCCPUSeconds:    gcCPU,
		SchedLatencyP99: p99,
		TickCounts:      p.snapshotTickCounts(),
	}
}

func (p *Probe) appendSample(s Sample) {
	p.ringMu.Lock()
	defer p.ringMu.Unlock()
	if len(p.ring) >= p.ringCap {
		copy(p.ring, p.ring[1:])
		p.ring = p.ring[:p.ringCap-1]
	}
	p.ring = append(p.ring, s)
}

func (p *Probe) snapshotSamples() []Sample {
	p.ringMu.Lock()
	defer p.ringMu.Unlock()
	out := make([]Sample, len(p.ring))
	copy(out, p.ring)
	return out
}

func (p *Probe) snapshotTickCounts() map[string]uint64 {
	p.tickMu.Lock()
	defer p.tickMu.Unlock()
	if len(p.tickCounts) == 0 {
		return nil
	}
	out := make(map[string]uint64, len(p.tickCounts))
	for k, v := range p.tickCounts {
		out[k] = v
	}
	return out
}

func histP99(h *metrics.Float64Histogram) float64 {
	var total uint64
	for _, c := range h.Counts {
		total += c
	}
	if total == 0 {
		return 0
	}
	target := uint64(float64(total) * 0.99)
	var acc uint64
	for i, c := range h.Counts {
		acc += c
		if acc >= target {
			return h.Buckets[i+1]
		}
	}
	return h.Buckets[len(h.Buckets)-1]
}
