// Unexported shows that reflection can read an unexported struct field but
// not write one, and that an exported field nested inside an unexported one
// is no more settable. The restriction does not propagate through an
// unexported *embedded* field, which is how promoted exported fields stay
// settable. See Section 12.5.
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"

	// os.File embeds an unexported *os.file; its layout is not part of
	// any promise and differs between releases and platforms.
	file := stdout.FieldByName("file").Elem()
	name := file.FieldByName("name")
	fmt.Println(name.String())                 // "/dev/stdout"
	fmt.Println(name.CanAddr(), name.CanSet()) // "true false"

	defer func() { fmt.Println("recovered:", recover()) }()
	name.SetString("/dev/null") // panic: unexported field
}
