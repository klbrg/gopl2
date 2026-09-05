// Hello prints a proverb chosen by the rsc.io/quote module.
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
