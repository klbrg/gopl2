// Errdiscrim shows the three ways to discriminate a file error:
// the type assertion, the os predicate, and errors.Is/errors.As.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)

	// The mechanism: a type assertion recovers the structured value.
	if pe, ok := err.(*fs.PathError); ok {
		fmt.Printf("%s %s: %v\n", pe.Op, pe.Path, pe.Err)
	}

	// The old predicate still works on an unwrapped os error.
	fmt.Println(os.IsNotExist(err))

	// The idiom, and the only one of the three that survives wrapping.
	fmt.Println(errors.Is(err, fs.ErrNotExist))

	wrapped := fmt.Errorf("loading config: %w", err)
	fmt.Println(os.IsNotExist(wrapped))
	fmt.Println(errors.Is(wrapped, fs.ErrNotExist))

	if pe, ok := errors.AsType[*fs.PathError](wrapped); ok {
		fmt.Println(pe.Path)
	}
}
