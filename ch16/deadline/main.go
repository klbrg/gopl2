// Deadline shows timeouts, deadlines, and cancellation causes.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ErrSlowBackend is the cause reported when the backend misses its budget.
var ErrSlowBackend = errors.New("backend exceeded its budget")

// work simulates an operation taking d, and reports the context's fate.
func work(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// A plain timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if dl, ok := ctx.Deadline(); ok {
		fmt.Println("budget:", time.Until(dl).Round(10*time.Millisecond))
	}
	err := work(ctx, 100*time.Millisecond)
	fmt.Println(err, errors.Is(err, context.DeadlineExceeded))
	fmt.Println("cause:", context.Cause(ctx))

	// The same timeout, with a cause of our own.
	ctx2, cancel2 := context.WithTimeoutCause(
		context.Background(), 20*time.Millisecond, ErrSlowBackend)
	defer cancel2()
	err = work(ctx2, 100*time.Millisecond)
	fmt.Println(err, errors.Is(err, context.DeadlineExceeded))
	fmt.Println("cause:", context.Cause(ctx2), errors.Is(context.Cause(ctx2), ErrSlowBackend))

	// A child deadline can only shorten the parent's, never extend it.
	parent, cancel3 := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel3()
	child, cancel4 := context.WithTimeout(parent, time.Hour)
	defer cancel4()
	dl, _ := child.Deadline()
	fmt.Println("child budget:", time.Until(dl).Round(10*time.Millisecond))
}
