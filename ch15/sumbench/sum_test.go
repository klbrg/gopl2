package sumbench

import "testing"

var (
	floats = make([]float64, 4096)
	ifaces = make([]Adder, 4096)
	sink   float64
)

func init() {
	for i := range floats {
		floats[i] = float64(i)
		ifaces[i] = F(i)
	}
}

func BenchmarkGeneric(b *testing.B) {
	for b.Loop() {
		sink = SumGeneric(floats)
	}
}

func BenchmarkConcrete(b *testing.B) {
	for b.Loop() {
		sink = SumFloat64(floats)
	}
}

func BenchmarkIface(b *testing.B) {
	for b.Loop() {
		sink = SumIface(ifaces)
	}
}
