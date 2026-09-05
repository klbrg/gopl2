// Package atoi converts decimal strings to integers.
package atoi

import "fmt"

// Atoi converts a decimal string, with an optional sign, to an int.
// (Our first attempt.)
func Atoi(s string) (int, error) {
	digits := s
	neg := false
	if len(digits) > 0 && (digits[0] == '+' || digits[0] == '-') {
		neg = digits[0] == '-'
		digits = digits[1:]
	}
	n := 0
	for i := 0; i < len(digits); i++ {
		c := digits[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("atoi: %q: invalid syntax", s)
		}
		n = n*10 + int(c-'0')
	}
	if neg {
		n = -n
	}
	return n, nil
}
