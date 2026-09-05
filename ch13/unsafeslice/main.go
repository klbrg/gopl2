// Unsafeslice shows the slice and string constructors added to package
// unsafe in Go 1.17 and Go 1.20, and the slice-to-array conversions.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [4]int32{1, 2, 3, 4}

	// A slice over the first two elements, without naming a[0:2].
	s := unsafe.Slice(&a[0], 2)
	fmt.Println(s, len(s), cap(s)) // "[1 2] 2 2"

	// SliceData recovers the pointer that Slice consumed.
	fmt.Println(unsafe.SliceData(s) == &a[0]) // "true"

	// String and StringData deconstruct and reconstruct a string
	// without copying its bytes.
	str := "hello"
	fmt.Println(unsafe.String(unsafe.StringData(str), len(str))) // "hello"

	// Slice to array pointer (Go 1.17), and slice to array (Go 1.20).
	bs := []byte{'G', 'o', '!', '!'}
	pa := (*[4]byte)(bs)
	arr := [4]byte(bs)
	fmt.Println(string(pa[:]), string(arr[:])) // "Go!! Go!!"

	// The conversion panics if the slice is too short.
	defer func() { fmt.Println("recovered:", recover()) }()
	_ = (*[8]byte)(bs)
}
