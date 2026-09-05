// Seqfuncs defines a few adaptors that turn one sequence into another.
package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Filter returns the values of seq for which keep returns true.
func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if keep(v) && !yield(v) {
				return
			}
		}
	}
}

// Map returns the results of applying f to each value of seq.
func Map[V, W any](seq iter.Seq[V], f func(V) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Take returns the first n values of seq.
func Take[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if n <= 0 || !yield(v) {
				return
			}
			n--
		}
	}
}

// Count yields 0, 1, 2, ... without end.
func Count(yield func(int) bool) {
	for i := 0; yield(i); i++ {
	}
}

func main() {
	odd := func(x int) bool { return x%2 == 1 }
	square := func(x int) int { return x * x }
	seq := Take(Map(Filter(Count, odd), square), 6)
	fmt.Println(slices.Collect(seq))

	words := slices.Values(strings.Fields("the quick brown fox"))
	long := Filter(words, func(s string) bool { return len(s) > 3 })
	fmt.Println(slices.Collect(Map(long, strings.ToUpper)))
}
