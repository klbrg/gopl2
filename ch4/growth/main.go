// Growth reports each capacity the runtime chooses as a slice is appended to.
//
// The slice is a package-level variable so that it escapes: for a slice that
// provably does not escape, the compiler may allocate a small backing array on
// the stack and the first capacity reported is that array's, not the runtime's.
package main

import "fmt"

var xs []int

func main() {
	prev := 0
	for i := 0; i < 2000; i++ {
		xs = append(xs, i)
		if cap(xs) != prev {
			fmt.Printf("len=%d\tcap=%d\n", len(xs), cap(xs))
			prev = cap(xs)
		}
	}
}
