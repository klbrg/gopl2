// Control shows how break, return, and defer behave across an iterator call.
package main

import (
	"fmt"
	"iter"
)

// counter yields 1, 2, 3, ... and reports when it is done.
func counter(name string) iter.Seq[int] {
	return func(yield func(int) bool) {
		defer fmt.Println(name, "cleaned up")
		for i := 1; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func find(target int) int {
	for x := range counter("find") {
		if x == target {
			return x // returns from find, not just the loop
		}
	}
	panic("unreachable")
}

func main() {
outer:
	for x := range counter("outer") {
		for y := range counter("inner") {
			fmt.Println(x, y)
			if y == 2 {
				continue outer
			}
			if x == 2 {
				break outer
			}
		}
	}
	fmt.Println("found", find(3))
}
