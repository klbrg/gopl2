// Package retry calls a function repeatedly until it succeeds.
package retry

import (
	"context"
	"time"
)

// Do calls f until it returns nil or ctx is done, waiting delay
// before the second attempt and twice as long before each attempt
// after that.  It returns nil if f succeeded, or the error from
// ctx if the context was cancelled first.
func Do(ctx context.Context, delay time.Duration, f func() error) error {
	for {
		if err := f(); err == nil {
			return nil
		}
		timer := time.NewTimer(delay)
		select {
		case <-timer.C:
			delay *= 2
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
}
