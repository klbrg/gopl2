package config

import (
	"os"
	"path/filepath"
	"testing"
)

// writeConf writes text to a fresh settings file and points
// GOPLCONF at it for the duration of the test.
func writeConf(t *testing.T, text string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "gopl.conf")
	if err := os.WriteFile(path, []byte(text), 0666); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GOPLCONF", path)
}

func TestLoad(t *testing.T) {
	writeConf(t, "# a comment\nname = gopher\n\nage=13\n")
	settings, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}
	for _, want := range []struct{ key, value string }{
		{"name", "gopher"},
		{"age", "13"},
	} {
		if got := settings[want.key]; got != want.value {
			t.Errorf("settings[%q] = %q, want %q",
				want.key, got, want.value)
		}
	}
	if len(settings) != 2 {
		t.Errorf("got %d settings, want 2", len(settings))
	}
}

func TestLoadSyntaxError(t *testing.T) {
	writeConf(t, "name = gopher\nnonsense\n")
	if _, err := Load(); err == nil {
		t.Error("Load() succeeded on a malformed file")
	}
}

func TestPathDefault(t *testing.T) {
	t.Setenv("GOPLCONF", "")
	if got, want := Path(), "gopl.conf"; got != want {
		t.Errorf("Path() = %q, want %q", got, want)
	}
}
