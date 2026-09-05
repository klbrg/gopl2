// Optional encodes an optional JSON field using new(expression), new in Go 1.26.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  *int   `json:"age,omitempty"` // age if known; nil otherwise
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / (365.25 * 24)) // approximately
}

func main() {
	born := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	newborn := time.Now()
	for _, p := range []Person{
		{Name: "alice", Age: new(yearsSince(born))},
		{Name: "bob", Age: new(yearsSince(newborn))},
		{Name: "carol"},
	} {
		data, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", data)
	}
}
