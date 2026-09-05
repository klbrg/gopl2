// The devnull command prints the path of the system's null device.
package main

import (
	"fmt"

	"github.com/klbrg/gopl2/ch10/nulldev"
)

func main() {
	fmt.Println(nulldev.Device)
}
