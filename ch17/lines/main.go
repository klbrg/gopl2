// Lines shows an iterator that reports failure through a Seq2 of value and error.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"
)

// Lines returns an iterator over the lines of r, paired with any read error.
// Iteration stops after the first error.
func Lines(r io.Reader) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			if !yield(sc.Text(), nil) {
				return
			}
		}
		if err := sc.Err(); err != nil {
			yield("", err)
		}
	}
}

// errReader fails after producing one complete line.
type errReader struct{ n int }

func (e *errReader) Read(p []byte) (int, error) {
	if e.n == 0 {
		e.n++
		return copy(p, "alpha\n"), nil
	}
	return 0, errors.New("disk on fire")
}

func main() {
	for line, err := range Lines(strings.NewReader("one\ntwo\nthree\n")) {
		fmt.Println(err, line)
	}

	for line, err := range Lines(&errReader{}) {
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		fmt.Println("line:", line)
	}
}
