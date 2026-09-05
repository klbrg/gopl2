// Inference demonstrates what type inference can and cannot work out.
package main

import (
	"fmt"
	"slices"
)

// !+decls
type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

// Scale multiplies each element of s by c.
// S is inferred from the argument; E follows from S's constraint.
func Scale[S ~[]E, E Number](s S, c E) S {
	out := make(S, len(s))
	for i, v := range s {
		out[i] = v * c
	}
	return out
}

// Map applies f to each element of s.
func Map[S ~[]E, E, R any](s S, f func(E) R) []R {
	out := make([]R, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

// double and isNegative are themselves generic.
func double[T Number](v T) T { return v + v }

func isNegative[T Number](v T) bool { return v < 0 }

// Zero returns the zero value of T. Its type parameter appears only
// in the result, so it can never be inferred.
func Zero[T any]() T {
	var z T
	return z
}

// !+celsius
type Celsius float64

type Readings []Celsius

func main() {
	r := Readings{-40, 0, 100}
	fmt.Printf("%T %v\n", Scale(r, 2), Scale(r, 2)) // "main.Readings [-80 0 200]"

	// Go 1.21: double is generic and is not instantiated here.
	fmt.Println(Map(r, double)) // "[-80 0 200]"

	// Go 1.21: a generic function assigned to a function-typed variable.
	var half func(float64) float64 = double
	fmt.Println(half(0.25)) // "0.5"

	// Go 1.21: a generic function passed to another generic function.
	fmt.Println(slices.IndexFunc(r, isNegative)) // "0"

	// Zero's type argument must be written out.
	fmt.Printf("%q %v\n", Zero[string](), Zero[Celsius]()) // `"" 0`
}
