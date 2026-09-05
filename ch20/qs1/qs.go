// Package qs escapes and unescapes the values of URL query parameters.
package qs

import (
	"fmt"
	"strconv"
	"strings"
)

const hexdigits = "0123456789abcdef"

// Escape returns s with every byte that is not an unreserved
// character replaced by "%" and two hexadecimal digits.
func Escape(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if c := s[i]; unreserved(c) {
			b.WriteByte(c)
		} else {
			b.WriteByte('%')
			b.WriteByte(hexdigits[c>>4])
			b.WriteByte(hexdigits[c&0xf])
		}
	}
	return b.String()
}

// unreserved reports whether c may appear in a query value as itself.
func unreserved(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' ||
		c == '-' || c == '.' || c == '_' || c == '~'
}

// Unescape reverses Escape.
func Unescape(s string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '%' {
			b.WriteByte(s[i])
			continue
		}
		n, err := strconv.ParseUint(s[i+1:i+3], 16, 8)
		if err != nil {
			return "", fmt.Errorf("qs: invalid escape %q", s[i:i+3])
		}
		b.WriteByte(byte(n))
		i += 2
	}
	return b.String(), nil
}
