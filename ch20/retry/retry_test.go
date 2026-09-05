package retry_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/klbrg/gopl2/ch20/retry"
)

func TestDoBackoff(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		start := time.Now()
		attempts := 0
		err := retry.Do(ctx, time.Second, func() error {
			attempts++
			if attempts < 4 {
				return errors.New("not yet")
			}
			return nil
		})
		if err != nil {
			t.Fatalf("Do() = %v, want nil", err)
		}
		if attempts != 4 {
			t.Errorf("f called %d times, want 4", attempts)
		}
		// 1s + 2s + 4s of backoff, none of it real.
		if got, want := time.Since(start), 7*time.Second; got != want {
			t.Errorf("elapsed = %v, want %v", got, want)
		}
	})
}

func TestDoCancellation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
		defer cancel()

		start := time.Now()
		err := retry.Do(ctx, time.Second, func() error {
			return errors.New("always fails")
		})
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Do() = %v, want DeadlineExceeded", err)
		}
		if got, want := time.Since(start), 10*time.Second; got != want {
			t.Errorf("elapsed = %v, want %v", got, want)
		}
	})
}

func TestDoProgress(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		// attempts is written by the retry goroutine and read by
		// this one, so it must be synchronized even in a bubble.
		var attempts atomic.Int64
		go retry.Do(ctx, time.Second, func() error {
			attempts.Add(1)
			return errors.New("always fails")
		})

		// Wait until the retry goroutine is durably blocked in
		// its timer, then check what it has done so far.
		start := time.Now()
		for _, want := range []struct {
			sleep time.Duration
			calls int64
		}{{0, 1}, {time.Second, 2}, {2 * time.Second, 3}} {
			time.Sleep(want.sleep)
			synctest.Wait()
			if got := attempts.Load(); got != want.calls {
				t.Fatalf("%v after start: %d attempts, want %d",
					time.Since(start), got, want.calls)
			}
		}
	})
}

// TestDoBackoffReal runs the same code outside a bubble, where
// the delays are real and the elapsed time can only be bounded.
func TestDoBackoffReal(t *testing.T) {
	start := time.Now()
	attempts := 0
	err := retry.Do(t.Context(), 100*time.Millisecond, func() error {
		attempts++
		if attempts < 4 {
			return errors.New("not yet")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Do() = %v, want nil", err)
	}
	t.Logf("elapsed %v", time.Since(start))
}
