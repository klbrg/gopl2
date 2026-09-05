// Ismethod gives an error type its own Is method.
package main

import (
	"errors"
	"fmt"
)

// ErrServerFault matches any 5xx status.
var ErrServerFault = errors.New("server fault")

// A StatusError is an HTTP status treated as an error.
type StatusError int

func (e StatusError) Error() string { return fmt.Sprintf("http status %d", int(e)) }

// Is reports whether e should be treated as equivalent to target.
func (e StatusError) Is(target error) bool {
	return target == ErrServerFault && e >= 500 && e < 600
}

func fetch(code int) error {
	return fmt.Errorf("fetch /orders: %w", StatusError(code))
}

func main() {
	for _, code := range []int{200, 404, 503} {
		err := fetch(code)
		fmt.Println(err, errors.Is(err, ErrServerFault))
	}
	// Equality still distinguishes individual statuses.
	fmt.Println(errors.Is(fetch(404), StatusError(404)))
}
