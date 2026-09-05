//go:build failing

// This test is expected to fail; see Section 19.7.
// Run it with: go test -tags failing ./stringfs1
package stringfs

import (
	"testing"
	"testing/fstest"
)

func TestStringFS(t *testing.T) {
	fsys := FS{
		"hello.txt":     "hello\n",
		"docs/read.txt": "read me\n",
	}
	if err := fstest.TestFS(fsys, "hello.txt", "docs/read.txt"); err != nil {
		t.Fatal(err)
	}
}
