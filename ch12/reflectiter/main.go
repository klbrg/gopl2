// Reflectiter shows the reflect additions this chapter mentions:
// TypeFor (Go 1.22), Value.Fields and Value.Methods (Go 1.26).
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

type Movie struct {
	Title, Subtitle string
	Year            int
}

// printMethods is the Print of Section 12.8, written with Value.Methods.
func printMethods(x any) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for m, mv := range v.Methods() {
		fmt.Printf("func (%s) %s%s\n", t, m.Name,
			strings.TrimPrefix(mv.Type().String(), "func"))
	}
}

func main() {
	// TypeFor names a type known at compile time, interfaces included.
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))            // "*os.File"
	fmt.Println(reflect.TypeFor[io.Writer]()) // "io.Writer"
	fmt.Println(reflect.TypeOf(w).Implements(
		reflect.TypeFor[io.Writer]())) // "true"

	// Value.Fields replaces the NumField/Field(i) loop.
	v := reflect.ValueOf(Movie{"Dr. Strangelove", "", 1964})
	for f, fv := range v.Fields() {
		fmt.Printf("%s = %v\n", f.Name, fv)
	}

	printMethods(time.Hour)
}
