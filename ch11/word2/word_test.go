package word

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input string
	want  bool
}{
	{"", true},
	{"a", true},
	{"aa", true},
	{"ab", false},
	{"kayak", true},
	{"detartrated", true},
	{"A man, a plan, a canal: Panama", true},
	{"Evil I did dwell; lewd did I live.", true},
	{"Able was I ere I saw Elba", true},
	{"été", true},
	{"Et se resservir, ivresse reste.", true},
	{"palindrome", false}, // non-palindrome
	{"desserts", false},   // semi-palindrome
}

// TestIsPalindrome runs each row of the table as a named subtest, so
// that a single row can be selected with -run=TestIsPalindrome/kayak.
func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if got := IsPalindrome(test.input); got != test.want {
				t.Errorf("IsPalindrome(%q) = %v", test.input, got)
			}
		})
	}
}

// BenchmarkIsPalindromeBN is the b.N idiom of section 11.4.
func BenchmarkIsPalindromeBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// BenchmarkIsPalindrome is the same benchmark written with b.Loop,
// which is the form to prefer since Go 1.24.
func BenchmarkIsPalindrome(b *testing.B) {
	for b.Loop() {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
