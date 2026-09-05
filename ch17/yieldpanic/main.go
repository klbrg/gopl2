// Yieldpanic shows what happens when an iterator ignores a false from yield.
package main

import "fmt"

// bad ignores yield's result and keeps going after the loop has stopped.
func bad(yield func(int) bool) {
	for i := range 3 {
		yield(i) // wrong: result discarded
	}
}

func main() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	for x := range bad {
		fmt.Println(x)
		break
	}
	fmt.Println("unreachable")
}
