package popcount

import "testing"

const x = 0x1234567890ABCDEF

// The old idiom: an explicit loop to b.N.
func BenchmarkPopCountBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

// The same benchmark written with b.Loop.
func BenchmarkPopCountLoop(b *testing.B) {
	for b.Loop() {
		PopCount(x)
	}
}

// The old workaround: assign to a package-level sink so that the
// compiler cannot discard the call.
var sink int

func BenchmarkPopCountSink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = PopCount(x)
	}
}
