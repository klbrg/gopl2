// Maxof demonstrates a single generic function that replaces one
// hand-written Max for every ordered type.
package main

import (
	"cmp"
	"fmt"
)

// !+max
// Max returns the larger of x and y.
func Max[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// !+celsius
type Celsius float64

func main() {
	fmt.Println(Max(3, 4))                  // "4"
	fmt.Println(Max("hello", "world"))      // "world"
	fmt.Println(Max(2.5, 2.25))             // "2.5"
	fmt.Println(Max(Celsius(-273.15), 100)) // "100"
	fmt.Println(Max[float64](1, 2))         // "2"
}
