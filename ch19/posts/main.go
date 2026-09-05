// posts lists the title of every article in a file system of Markdown files.
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

//go:embed posts
var content embed.FS

// titles returns the first heading of every Markdown file in fsys.
func titles(fsys fs.FS) ([]string, error) {
	names, err := fs.Glob(fsys, "*.md")
	if err != nil {
		return nil, err
	}
	var out []string
	for _, name := range names {
		data, err := fs.ReadFile(fsys, name)
		if err != nil {
			return nil, err
		}
		line, _, _ := strings.Cut(string(data), "\n")
		out = append(out, strings.TrimPrefix(line, "# "))
	}
	return out, nil
}

func main() {
	embedded, err := fs.Sub(content, "posts")
	if err != nil {
		log.Fatal(err)
	}
	for _, fsys := range []fs.FS{embedded, os.DirFS("posts/posts")} {
		names, err := titles(fsys)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(names)
	}
}
