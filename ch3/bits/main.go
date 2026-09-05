// Bits shows a uint8 used as a set of 8 bits, and the same questions
// answered by the math/bits package.
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := range 8 {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	fmt.Println(bits.OnesCount8(x))        // "2", the size of the set {1, 5}
	fmt.Println(bits.TrailingZeros8(x))    // "1", the smallest member
	fmt.Println(bits.Len8(x))              // "6", one more than the largest
	fmt.Printf("%08b\n", bits.Reverse8(x)) // "01000100", the bits end for end

	// The empty set: TrailingZeros reports the width, not zero.
	fmt.Println(bits.OnesCount8(0), bits.Len8(0), bits.TrailingZeros8(0)) // "0 0 8"
}
