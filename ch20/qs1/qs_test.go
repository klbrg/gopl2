package qs

import "testing"

func TestEscape(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"", ""},
		{"gopher", "gopher"},
		{"a=1&b=2", "a%3d1%26b%3d2"},
		{"100%", "100%25"},
		{"été", "%c3%a9t%c3%a9"},
	}
	for _, test := range tests {
		if got := Escape(test.input); got != test.want {
			t.Errorf("Escape(%q) = %q, want %q",
				test.input, got, test.want)
		}
	}
}

// FuzzRoundTrip checks that Unescape undoes Escape.
func FuzzRoundTrip(f *testing.F) {
	f.Add("")
	f.Add("gopher")
	f.Add("a=1&b=2")
	f.Add("100%")
	f.Add("été")
	f.Fuzz(func(t *testing.T, s string) {
		out, err := Unescape(Escape(s))
		if err != nil {
			t.Fatalf("Unescape(Escape(%q)) failed: %v", s, err)
		}
		if out != s {
			t.Errorf("Unescape(Escape(%q)) = %q", s, out)
		}
	})
}

// FuzzUnescape checks that Unescape either rejects a string or
// decodes it to something Escape maps back to the same value.
func FuzzUnescape(f *testing.F) {
	f.Add("gopher")
	f.Add("a%3d1")
	f.Add("%zz")
	f.Fuzz(func(t *testing.T, s string) {
		out, err := Unescape(s)
		if err != nil {
			return // rejecting malformed input is not a bug
		}
		again, err := Unescape(Escape(out))
		if err != nil {
			t.Fatalf("Unescape(Escape(%q)) failed: %v", out, err)
		}
		if again != out {
			t.Errorf("Unescape(%q) = %q, but round trip gave %q",
				s, out, again)
		}
	})
}
