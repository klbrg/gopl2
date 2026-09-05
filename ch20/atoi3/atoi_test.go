package atoi

import (
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"0", 0},
		{"7", 7},
		{"+7", 7},
		{"-7", -7},
		{"1729", 1729},
	}
	for _, test := range tests {
		got, err := Atoi(test.input)
		if err != nil {
			t.Errorf("Atoi(%q) failed: %v", test.input, err)
			continue
		}
		if got != test.want {
			t.Errorf("Atoi(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

// FuzzAtoi compares Atoi against strconv.Atoi, which we trust.
func FuzzAtoi(f *testing.F) {
	f.Add("0")
	f.Add("-1")
	f.Add("1729")
	f.Add("x")
	f.Fuzz(func(t *testing.T, s string) {
		got, gotErr := Atoi(s)
		want, wantErr := strconv.Atoi(s)
		if (gotErr != nil) != (wantErr != nil) {
			t.Fatalf("Atoi(%q) error = %v; strconv.Atoi error = %v",
				s, gotErr, wantErr)
		}
		if gotErr == nil && got != want {
			t.Errorf("Atoi(%q) = %d; strconv.Atoi = %d", s, got, want)
		}
	})
}
