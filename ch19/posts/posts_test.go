package main

import (
	"slices"
	"testing"
	"testing/fstest"
)

func TestTitles(t *testing.T) {
	fsys := fstest.MapFS{
		"b.md":      {Data: []byte("# Beta\nbody\n")},
		"a.md":      {Data: []byte("# Alpha\nbody\n")},
		"notes.txt": {Data: []byte("# Ignored\n")},
		"sub/c.md":  {Data: []byte("# Gamma\n")},
		"empty.md":  {Data: []byte("")},
	}
	got, err := titles(fsys)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"Alpha", "Beta", ""}
	if !slices.Equal(got, want) {
		t.Errorf("titles() = %q, want %q", got, want)
	}
}
