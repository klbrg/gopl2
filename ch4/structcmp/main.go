// Structcmp shows that struct comparison stops at the first mismatched field.
package main

import "fmt"

type pair struct {
	tag  string
	data any
}

func main() {
	a := pair{"x", []int{1}}
	b := pair{"y", []int{2}}
	c := pair{"x", []int{3}}

	fmt.Println(a == b) // "false"
	fmt.Println(a == c) // panic: comparing uncomparable type []int
}
