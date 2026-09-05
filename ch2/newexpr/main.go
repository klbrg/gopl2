// Newexpr shows the expression form of the built-in new function that
// Go 1.26 added alongside new(T).  It backs the listings in Section 2.3.3.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Person has an optional Age: a nil pointer means the age is unknown.
type Person struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / (365.25 * 24)) // approximately
}

func main() {
	p := new(2) // *int, pointing to an int variable holding 2
	fmt.Println(*p)
	s := new("hello") // *string
	fmt.Println(*s)
	fmt.Printf("%T %T %T\n", new(0), new(0.0), new('a'))

	born := time.Date(1994, time.March, 1, 0, 0, 0, 0, time.UTC)
	alice := Person{Name: "Alice", Age: new(yearsSince(born))}
	bob := Person{Name: "Bob"}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(alice)
	enc.Encode(bob)
}
