// logfs shows how fs.ReadFile and fs.ReadDir fall back to Open, and what
// a wrapper that implements only Open costs.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"testing/fstest"
)

// A logFS wraps a file system and reports every call to Open.
type logFS struct{ fsys fs.FS }

func (l logFS) Open(name string) (fs.File, error) {
	fmt.Printf("\topen %q\n", name)
	return l.fsys.Open(name)
}

func main() {
	base := fstest.MapFS{
		"docs/a.txt": {Data: []byte("alpha\n")},
		"docs/b.txt": {Data: []byte("beta\n")},
	}

	_, isReadFileFS := any(base).(fs.ReadFileFS)
	_, wrapperIsReadFileFS := any(logFS{base}).(fs.ReadFileFS)
	fmt.Println("MapFS implements fs.ReadFileFS:", isReadFileFS)
	fmt.Println("logFS implements fs.ReadFileFS:", wrapperIsReadFileFS)

	fmt.Println("fs.ReadFile on the wrapper:")
	data, err := fs.ReadFile(logFS{base}, "docs/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\tread %d bytes\n", len(data))

	fmt.Println("fs.ReadDir on the wrapper:")
	entries, err := fs.ReadDir(logFS{base}, "docs")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\tread %d entries\n", len(entries))

	fmt.Println("fs.ReadFile on the MapFS itself:")
	if _, err := fs.ReadFile(base, "docs/a.txt"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("\t(no Open: MapFS.ReadFile was called directly)")
}
