// Adder demonstrates a self-referential type constraint,
// permitted since Go 1.26.
package main

import "fmt"

// !+decls
// An Adder is a type that can add another value of its own type
// to itself. The constraint refers to the type it constrains.
type Adder[A Adder[A]] interface {
	Add(A) A
}

// Sum adds the values, which must all have the same type.
func Sum[A Adder[A]](first A, rest ...A) A {
	total := first
	for _, v := range rest {
		total = total.Add(v)
	}
	return total
}

// !+types
type Vec2 struct{ X, Y float64 }

func (v Vec2) Add(w Vec2) Vec2 { return Vec2{v.X + w.X, v.Y + w.Y} }

type Money int64 // in cents

func (m Money) Add(n Money) Money { return m + n }

func main() {
	fmt.Println(Sum(Vec2{1, 2}, Vec2{3, 4}, Vec2{5, 6})) // "{9 12}"
	fmt.Println(Sum(Money(199), Money(250)))             // "449"
}
