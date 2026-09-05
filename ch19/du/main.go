// du reports the number and total size of the files beneath each
// directory named on the command line.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

var skip = flag.String("skip", ".git", "name of directories to skip")

// diskUsage counts the files in the subtree of fsys rooted at root.
func diskUsage(fsys fs.FS, root string) (nfiles, nbytes int64, err error) {
	err = fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if d == nil {
				return err // the root itself could not be read
			}
			fmt.Fprintf(os.Stderr, "du: %v\n", err)
			return nil // report and continue
		}
		if d.IsDir() {
			if d.Name() == *skip && path != root {
				return fs.SkipDir
			}
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		nfiles++
		nbytes += info.Size()
		return nil
	})
	return
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	for _, root := range roots {
		nfiles, nbytes, err := diskUsage(os.DirFS(root), ".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "du: %v\n", err)
			continue
		}
		fmt.Printf("%s: %d files  %.1f KB\n", root, nfiles, float64(nbytes)/1e3)
	}
}
