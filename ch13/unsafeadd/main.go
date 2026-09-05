// Unsafeadd is the unsafe.Add form of the pointer arithmetic in Section 13.2.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Add(unsafe.Pointer(&x), unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
}
