// Advice collects the small examples from the section on when not to
// use type parameters.
package main

import (
	"fmt"
	"io"
	"strings"
)

// !+reader
// ReadAllG is a needless generalization: the type parameter R adds
// nothing that the io.Reader interface does not already provide, and
// it forces a fresh instantiation for every argument type.
func ReadAllG[R io.Reader](r R) ([]byte, error) { return io.ReadAll(r) }

// ReadAll is the version to write.
func ReadAll(r io.Reader) ([]byte, error) { return io.ReadAll(r) }

// !+heterogeneous
// A []Shape can hold a Circle and a Square at the same time.
// A []T, for a type parameter T, cannot: every element has one type.
type Shape interface{ Area() float64 }

type Circle struct{ R float64 }

func (c Circle) Area() float64 { return 3.14159 * c.R * c.R }

type Square struct{ S float64 }

func (s Square) Area() float64 { return s.S * s.S }

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	b, _ := ReadAll(strings.NewReader("hello"))
	fmt.Printf("%s\n", b) // "hello"

	fmt.Println(TotalArea([]Shape{Circle{1}, Square{2}})) // "7.14159"
	_ = ReadAllG[*strings.Reader]
}
