// Metrics reports a selection of the Go runtime's own metrics.
package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/metrics"

	"github.com/klbrg/gopl2/ch18/notime"
)

var names = []string{
	"/sched/goroutines:goroutines",
	"/memory/classes/heap/objects:bytes",
	"/gc/heap/allocs:bytes",
	"/gc/cycles/total:gc-cycles",
	"/sched/latencies:seconds",
}

func main() {
	fmt.Printf("%d metrics are supported by this runtime\n", len(metrics.All()))

	samples := make([]metrics.Sample, len(names))
	for i, name := range names {
		samples[i].Name = name
	}
	churn()
	metrics.Read(samples)

	logger := slog.New(slog.NewTextHandler(os.Stdout, notime.Options(nil)))
	for _, s := range samples {
		switch s.Value.Kind() {
		case metrics.KindUint64:
			logger.Info("metric", "name", s.Name, "value", s.Value.Uint64())
		case metrics.KindFloat64:
			logger.Info("metric", "name", s.Name, "value", s.Value.Float64())
		case metrics.KindFloat64Histogram:
			h := s.Value.Float64Histogram()
			logger.Info("metric", "name", s.Name,
				"count", total(h), "median", quantile(h, 0.5),
				"p99", quantile(h, 0.99))
		case metrics.KindBad:
			logger.Error("unknown metric", "name", s.Name)
		}
	}
}

// churn allocates and discards memory so that the metrics are not all zero.
func churn() {
	var sink [][]byte
	for i := 0; i < 1000; i++ {
		sink = append(sink, make([]byte, 4096))
		if len(sink) == 100 {
			sink = nil
		}
	}
}

func total(h *metrics.Float64Histogram) uint64 {
	var n uint64
	for _, c := range h.Counts {
		n += c
	}
	return n
}

// quantile returns the upper bound of the bucket containing the
// q-quantile of h.
func quantile(h *metrics.Float64Histogram, q float64) float64 {
	want := uint64(q * float64(total(h)))
	var n uint64
	for i, c := range h.Counts {
		if n += c; n >= want {
			return h.Buckets[i+1]
		}
	}
	return h.Buckets[len(h.Buckets)-1]
}
