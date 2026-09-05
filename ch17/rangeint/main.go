// Rangeint demonstrates the range-over-integer form added in Go 1.22.
package main

import "fmt"

func main() {
	for i := range 5 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// The loop variable may be omitted entirely.
	n := 0
	for range 3 {
		n++
	}
	fmt.Println("n =", n)

	// An integer-typed operand fixes the type of the iteration value.
	var limit int8 = 4
	for i := range limit {
		fmt.Printf("%T:%d ", i, i)
	}
	fmt.Println()

	// A non-positive operand runs zero iterations.
	for i := range -1 {
		fmt.Println("unreachable", i)
	}

	// Countdown, the old way and the new.
	for i := 3; i > 0; i-- {
		fmt.Print(i, " ")
	}
	for i := range 3 {
		fmt.Print(3-i, " ")
	}
	fmt.Println()
}
