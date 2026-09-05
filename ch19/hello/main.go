// hello prints file contents embedded into the binary at compile time.
package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed hello.txt
var greeting string

//go:embed hello.txt
var raw []byte

//go:embed hello.txt quote.txt
var files embed.FS

func main() {
	fmt.Print(greeting)
	fmt.Println(len(raw), "bytes")

	data, err := files.ReadFile("quote.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))

	entries, err := files.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %d\n", e.Name(), info.Size())
	}
}
