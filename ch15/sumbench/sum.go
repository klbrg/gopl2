// Package sumbench compares three ways of writing the same summation.
package sumbench

// !+
type Number interface{ ~int | ~float64 }

// Generic: one source, instantiated per type argument.
func SumGeneric[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// Concrete: written out by hand for float64.
func SumFloat64(s []float64) float64 {
	var total float64
	for _, v := range s {
		total += v
	}
	return total
}

// Interface-based: one implementation, dynamic dispatch per element.
type Adder interface{ Add(float64) float64 }

func SumIface(s []Adder) float64 {
	var total float64
	for _, v := range s {
		total = v.Add(total)
	}
	return total
}

type F float64

func (f F) Add(x float64) float64 { return x + float64(f) }
