// Runeconv shows integer-to-string conversion written so that go vet's
// stringintconv check is satisfied.
package main

import "fmt"

func main() {
	fmt.Println(string(rune(65)))      // "A", not "65"
	fmt.Println(string(rune(0x4eac)))  // "京"
	fmt.Println(string(rune(1234567))) // "�", the replacement character

	// The conversion the check is warning about: this is strconv.Itoa's job.
	// fmt.Println(string(65)) // vet: yields a string of one rune
}
