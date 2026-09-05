// Package word provides utilities for word games.
//
// This is the word2 version from chapter 11, kept here so that the
// listings the second edition adds to that chapter -- the subtest form
// of the table-driven test and the b.Loop form of the benchmark -- can
// be compiled and run.
package word

import "unicode"

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
