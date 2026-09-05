// Join reports every problem with a record, not just the first.
package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrMissing = errors.New("missing")
	ErrTooLong = errors.New("too long")
)

// A User is a record awaiting validation.
type User struct {
	Name  string
	Email string
	Bio   string
}

// Validate reports every problem it finds with u.
func (u *User) Validate() error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, fmt.Errorf("name: %w", ErrMissing))
	}
	if !strings.Contains(u.Email, "@") {
		errs = append(errs, fmt.Errorf("email %q: %w", u.Email, ErrMissing))
	}
	if len(u.Bio) > 16 {
		errs = append(errs, fmt.Errorf("bio (%d bytes): %w", len(u.Bio), ErrTooLong))
	}
	return errors.Join(errs...)
}

func main() {
	u := &User{Email: "gopher", Bio: strings.Repeat("x", 20)}
	err := u.Validate()
	fmt.Println(err)

	fmt.Println(errors.Is(err, ErrMissing), errors.Is(err, ErrTooLong))
	fmt.Println(errors.Unwrap(err)) // Unwrap does not descend into a Join

	var multi interface{ Unwrap() []error }
	if errors.As(err, &multi) {
		fmt.Println(len(multi.Unwrap()), "problems")
	}

	good := &User{Name: "Alice", Email: "alice@example.com"}
	fmt.Println(good.Validate() == nil)
}
