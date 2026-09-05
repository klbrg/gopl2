// Strictly demonstrates the comparable constraint and the difference
// between comparable and strictly comparable types.
package main

import "fmt"

// !+index
// Index returns the index of the first occurrence of v in s, or -1.
func Index[E comparable](s []E, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

// !+record
// A Record is comparable but not strictly comparable: the Payload
// field is an interface.
type Record struct {
	ID      int
	Payload any
}

func main() {
	fmt.Println(Index([]string{"a", "b"}, "b")) // "1"

	// any is comparable, though not strictly comparable.
	vals := []any{1, "two", 3.0}
	fmt.Println(Index(vals, "two")) // "1"

	recs := []Record{{1, "x"}, {2, "y"}}
	fmt.Println(Index(recs, Record{2, "y"})) // "1"

	defer func() {
		fmt.Println("recovered:", recover())
	}()
	var slice any = []int{1, 2}
	bad := []any{slice}
	fmt.Println(Index(bad, slice)) // panics: []int is not comparable
}
