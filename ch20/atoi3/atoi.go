// Package atoi converts decimal strings to integers.
package atoi

import (
	"fmt"
	"math"
)

// Atoi converts a decimal string, with an optional sign, to an int.
func Atoi(s string) (int, error) {
	digits := s
	neg := false
	if len(digits) > 0 && (digits[0] == '+' || digits[0] == '-') {
		neg = digits[0] == '-'
		digits = digits[1:]
	}
	if len(digits) == 0 {
		return 0, fmt.Errorf("atoi: %q: invalid syntax", s)
	}
	// The magnitude of the most negative int is one greater
	// than that of the most positive one.
	limit := uint(math.MaxInt)
	if neg {
		limit++
	}
	var n uint
	for i := 0; i < len(digits); i++ {
		c := digits[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("atoi: %q: invalid syntax", s)
		}
		d := uint(c - '0')
		if n > (limit-d)/10 {
			if neg {
				return math.MinInt, fmt.Errorf("atoi: %q: out of range", s)
			}
			return math.MaxInt, fmt.Errorf("atoi: %q: out of range", s)
		}
		n = n*10 + d
	}
	if neg {
		return -int(n), nil // -int(1<<63) is MinInt, as required
	}
	return int(n), nil
}
