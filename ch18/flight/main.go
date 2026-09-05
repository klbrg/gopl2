// Flight keeps a rolling window of execution trace data in memory and
// writes it out only when a request takes too long.
package main

import (
	"log/slog"
	"os"
	"runtime/trace"
	"time"
)

const slow = 20 * time.Millisecond

func main() {
	fr := trace.NewFlightRecorder(trace.FlightRecorderConfig{
		MinAge:   2 * time.Second,
		MaxBytes: 1 << 20,
	})
	if err := fr.Start(); err != nil {
		slog.Error("flight recorder", "err", err)
		os.Exit(1)
	}
	defer fr.Stop()

	var dumped bool
	for i := range 20 {
		d := handle(i)
		if d > slow && !dumped {
			dumped = true
			f, err := os.Create("flight.out")
			if err != nil {
				slog.Error("create", "err", err)
				continue
			}
			n, err := fr.WriteTo(f)
			f.Close()
			slog.Info("snapshot written", "request", i,
				"elapsed", d.Round(time.Millisecond), "bytes", n, "err", err)
		}
	}
}

// handle simulates a request that is occasionally slow.
func handle(i int) time.Duration {
	start := time.Now()
	d := time.Millisecond
	if i == 13 {
		d = 50 * time.Millisecond
	}
	time.Sleep(d)
	return time.Since(start)
}
