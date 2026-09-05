// Tracedemo writes an execution trace of a small pipeline to trace.out.
package main

import (
	"context"
	"crypto/sha256"
	"log/slog"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		slog.Error("create", "err", err)
		os.Exit(1)
	}
	defer f.Close()
	if err := trace.Start(f); err != nil {
		slog.Error("trace.Start", "err", err)
		os.Exit(1)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for id := range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			process(context.Background(), id)
		}()
	}
	wg.Wait()
}

// process handles one item, recording a task and two regions.
func process(ctx context.Context, id int) {
	ctx, task := trace.NewTask(ctx, "process")
	defer task.End()
	trace.Logf(ctx, "item", "id=%d", id)

	var data []byte
	trace.WithRegion(ctx, "read", func() { data = read(id) })
	trace.WithRegion(ctx, "digest", func() { digest(data) })
}

// read simulates a slow I/O operation.
func read(id int) []byte {
	time.Sleep(time.Duration(id) * time.Millisecond)
	return make([]byte, 1<<16)
}

// digest simulates a CPU-bound operation.
func digest(data []byte) [32]byte {
	sum := sha256.Sum256(data)
	for range 200 {
		sum = sha256.Sum256(sum[:])
	}
	return sum
}
