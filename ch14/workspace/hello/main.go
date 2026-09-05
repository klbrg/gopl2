package main

import (
	"fmt"

	"github.com/klbrg/gopl2/ch14/workspace/greeting"
	"rsc.io/quote"
)

func main() {
	fmt.Println(greeting.Hello("workspace"))
	fmt.Println(quote.Glass())
}
