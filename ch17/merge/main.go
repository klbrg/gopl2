// Merge interleaves two sorted sequences using pull-style iterators.
package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

// Merge returns the values of x and y, both sorted, as a single sorted sequence.
func Merge[E cmp.Ordered](x, y iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		nextX, stopX := iter.Pull(x)
		defer stopX()
		nextY, stopY := iter.Pull(y)
		defer stopY()

		vx, okx := nextX()
		vy, oky := nextY()
		for okx && oky {
			if vx <= vy {
				if !yield(vx) {
					return
				}
				vx, okx = nextX()
			} else {
				if !yield(vy) {
					return
				}
				vy, oky = nextY()
			}
		}
		for ; okx; vx, okx = nextX() {
			if !yield(vx) {
				return
			}
		}
		for ; oky; vy, oky = nextY() {
			if !yield(vy) {
				return
			}
		}
	}
}

func main() {
	a := slices.Values([]int{1, 4, 9, 16})
	b := slices.Values([]int{2, 3, 5, 8, 13})
	fmt.Println(slices.Collect(Merge(a, b)))

	// next keeps returning false once the sequence is exhausted.
	next, stop := iter.Pull(slices.Values([]string{"x"}))
	defer stop()
	for range 3 {
		v, ok := next()
		fmt.Printf("%q %v\n", v, ok)
	}
}
