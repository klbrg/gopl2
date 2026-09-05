//go:build go1.21

// Downgrade is loopvar with one line added: the build constraint above
// lowers this file's language version to Go 1.21, so the loop variable is
// created once and shared by every closure.
package main

import "fmt"

func main() {
	var seen []string
	var fns []func()
	for _, s := range []string{"a", "b", "c"} {
		fns = append(fns, func() { seen = append(seen, s) })
	}
	for _, f := range fns {
		f()
	}
	fmt.Println(seen)
}
