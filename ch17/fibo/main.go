// Fibo demonstrates ranging over a function, added in Go 1.23.
package main

import "fmt"

// fibo yields the Fibonacci sequence, which never ends.
func fibo(yield func(int) bool) {
	f0, f1 := 0, 1
	for yield(f0) {
		f0, f1 = f1, f0+f1
	}
}

// countdown yields nothing at all; it is a func(func() bool).
func countdown(yield func() bool) {
	for range 3 {
		if !yield() {
			return
		}
	}
}

// pairs yields index/value pairs, like a slice.
func pairs(yield func(int, string) bool) {
	for i, s := range []string{"a", "b", "c"} {
		if !yield(i, s) {
			return
		}
	}
}

func main() {
	for x := range fibo {
		if x >= 100 {
			break
		}
		fmt.Print(x, " ")
	}
	fmt.Println()

	ticks := 0
	for range countdown {
		ticks++
	}
	fmt.Println("ticks =", ticks)

	for i, s := range pairs {
		fmt.Printf("%d=%s ", i, s)
	}
	fmt.Println()

	// Only the first iteration value may be taken from a one-value iterator.
	for x := range fibo {
		fmt.Println("first =", x)
		break
	}
}
