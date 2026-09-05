// Firsterr runs tasks in parallel and cancels the rest when one fails.
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// A Task does some work, stopping early if ctx is done.
type Task func(ctx context.Context) error

// Run runs tasks concurrently. The first failure cancels the others
// and becomes the cause recorded in the derived context.
func Run(ctx context.Context, tasks ...Task) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	var wg sync.WaitGroup
	var once sync.Once
	var first error
	for i, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := task(ctx); err != nil {
				once.Do(func() {
					first = fmt.Errorf("task %d: %w", i, err)
					cancel(first)
				})
			}
		}()
	}
	wg.Wait()
	return first
}

// ErrRefused reports that a backend refused the connection.
var ErrRefused = errors.New("connection refused")

// sleeper returns a task that works for d and then fails with err.
func sleeper(d time.Duration, err error) Task {
	return func(ctx context.Context) error {
		select {
		case <-time.After(d):
			return err
		case <-ctx.Done():
			return context.Cause(ctx)
		}
	}
}

func main() {
	err := Run(context.Background(),
		sleeper(10*time.Millisecond, ErrRefused),
		sleeper(time.Hour, nil),
		sleeper(time.Hour, nil),
	)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrRefused))
}
