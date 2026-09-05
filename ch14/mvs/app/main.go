package main

import (
	"fmt"

	"github.com/klbrg/gopl2/ch14/mvs/a"
	"github.com/klbrg/gopl2/ch14/mvs/b"
)

func main() {
	fmt.Println("a sees", a.C())
	fmt.Println("b sees", b.C())
}
