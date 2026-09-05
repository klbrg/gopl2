// Twoversions imports two incompatible major versions of the same library.
package main

import (
	"fmt"

	"rsc.io/quote"
	quotev3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(quotev3.HelloV3())
}
