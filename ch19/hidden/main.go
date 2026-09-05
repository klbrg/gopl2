// hidden shows how //go:embed treats files whose names begin with '.' or '_'.
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed assets
var tree embed.FS

//go:embed assets/*
var star embed.FS

//go:embed all:assets
var all embed.FS

func list(name string, fsys fs.FS) {
	fmt.Println(name)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println("   ", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	list("assets", tree)
	list("assets/*", star)
	list("all:assets", all)
}
