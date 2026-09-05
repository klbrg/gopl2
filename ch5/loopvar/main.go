// Loopvar shows what a closure over a loop variable captures. The
// language version comes from the go directive in go.mod, so editing
// that one line changes what this program prints.
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
