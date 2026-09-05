package sum

import "testing"

// input builds a large slice and reports that it did so.
func input(b *testing.B) []int {
	b.Logf("building the input")
	xs := make([]int, 1e6)
	for i := range xs {
		xs[i] = i
	}
	return xs
}

var sink int

func BenchmarkSumBN(b *testing.B) {
	xs := input(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = Sum(xs)
	}
}

func BenchmarkSumLoop(b *testing.B) {
	xs := input(b)
	for b.Loop() {
		Sum(xs)
	}
}
