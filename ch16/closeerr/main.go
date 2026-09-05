// Closeerr reports the write error and the close error together.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// save writes data to w and closes it, reporting both failures if both occur.
func save(w io.WriteCloser, data []byte) (err error) {
	defer func() {
		if cerr := w.Close(); cerr != nil {
			err = errors.Join(err, fmt.Errorf("save: %w", cerr))
		}
	}()
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("save: %w", err)
	}
	return nil
}

// A brokenPipe fails on both write and close.
type brokenPipe struct{}

var ErrClosed = errors.New("connection reset")

func (brokenPipe) Write(p []byte) (int, error) {
	return 0, errors.New("no space left on device")
}
func (brokenPipe) Close() error { return ErrClosed }

func main() {
	dir, err := os.MkdirTemp("", "ch16")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "conf"))
	if err != nil {
		panic(err)
	}
	fmt.Println(save(f, []byte("port = 8000\n")))

	err = save(brokenPipe{}, []byte("port = 8000\n"))
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrClosed))
}
