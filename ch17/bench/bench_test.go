// Package bench measures the cost of the three ways to walk a slice.
package bench

import (
	"iter"
	"slices"
	"testing"
)

var data = make([]int, 1000)

var sink int

func BenchmarkDirect(b *testing.B) {
	for b.Loop() {
		sum := 0
		for _, v := range data {
			sum += v
		}
		sink = sum
	}
}

func BenchmarkValues(b *testing.B) {
	for b.Loop() {
		sum := 0
		for v := range slices.Values(data) {
			sum += v
		}
		sink = sum
	}
}

func BenchmarkPull(b *testing.B) {
	for b.Loop() {
		sum := 0
		next, stop := iter.Pull(slices.Values(data))
		for {
			v, ok := next()
			if !ok {
				break
			}
			sum += v
		}
		stop()
		sink = sum
	}
}
