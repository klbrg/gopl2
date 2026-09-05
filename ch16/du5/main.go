// Du5 computes the disk usage of the files in a directory,
// stopping early if its context is canceled.
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// sema is a counting semaphore limiting concurrent reads of directories.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(ctx context.Context, dir string) []os.DirEntry {
	select {
	case sema <- struct{}{}: // acquire token
	case <-ctx.Done():
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(ctx context.Context, dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if ctx.Err() != nil {
		return
	}
	for _, e := range dirents(ctx, dir) {
		if e.IsDir() {
			wg.Add(1)
			go walkDir(ctx, filepath.Join(dir, e.Name()), wg, fileSizes)
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		select {
		case fileSizes <- info.Size():
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(ctx, root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-ctx.Done():
			break loop // no drain: every sender selects on ctx.Done() too
		}
	}
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1e3)
	if err := ctx.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "du: %v (partial results)\n", err)
	}
}
