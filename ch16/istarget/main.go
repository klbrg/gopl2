// Istarget checks errors.Is with nil and with an uncomparable target.
package main

import (
	"errors"
	"fmt"
)

type badTarget []int

func (badTarget) Error() string { return "bad" }

func main() {
	fmt.Println(errors.Is(nil, nil))
	fmt.Println(errors.Is(errors.New("x"), nil))
	defer func() { fmt.Println("recovered:", recover()) }()
	fmt.Println(errors.Is(errors.New("x"), badTarget{1}))
}
