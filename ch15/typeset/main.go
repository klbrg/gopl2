// Typeset demonstrates constraint type sets and the ~ token.
package main

import "fmt"

// !+number
// A Number is any type whose underlying type is one of Go's
// numeric types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum returns the sum of the elements of s.
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// !+units
type Bytes int64
type Celsius float64

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))             // "6"
	fmt.Println(Sum([]float64{1.5, 2.5}))        // "4"
	fmt.Println(Sum([]Bytes{1024, 2048}))        // "3072"
	fmt.Println(Sum([]Celsius{-40, 100}))        // "60"
	fmt.Printf("%T %[1]v\n", Sum([]Bytes{1, 2})) // "main.Bytes 3"
}
