// Literals shows the numeric literal notations added in Go 1.13.
package main

import "fmt"

func main() {
	fmt.Println(0666, 0o666)  // "438 438"
	fmt.Println(0b1011)       // "11"
	fmt.Println(1_000_000)    // "1000000"
	fmt.Println(0x_dead_beef) // "3735928559"
	fmt.Println(0x1p-2)       // "0.25", a hexadecimal float
	fmt.Printf("%#o %#x\n", 438, 3735928559)
}
