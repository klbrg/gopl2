// Testsum reads the JSON event stream of "go test -json" on its
// standard input and reports the failures and the slowest tests.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type event struct {
	Action  string
	Package string
	Test    string
	Elapsed float64
	Output  string
}

type result struct {
	name    string
	elapsed float64
	output  []string
}

func main() {
	results := make(map[string]*result)
	var order []string
	dec := json.NewDecoder(os.Stdin)
	for {
		var e event
		if err := dec.Decode(&e); err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "testsum: %v\n", err)
			os.Exit(1)
		}
		if e.Test == "" {
			continue // a package-level event
		}
		key := e.Package + "." + e.Test
		r, ok := results[key]
		if !ok {
			r = &result{name: key}
			results[key], order = r, append(order, key)
		}
		switch e.Action {
		case "output":
			r.output = append(r.output, e.Output)
		case "pass", "fail", "skip":
			r.elapsed = e.Elapsed
			if e.Action != "fail" {
				r.output = nil // keep output only for failures
			}
		}
	}

	for _, key := range order {
		if r := results[key]; r.output != nil {
			fmt.Printf("FAIL %s\n", r.name)
			fmt.Print(strings.Join(r.output, ""))
		}
	}

	slowest := make([]*result, 0, len(results))
	for _, r := range results {
		slowest = append(slowest, r)
	}
	slices.SortFunc(slowest, func(a, b *result) int {
		switch {
		case a.elapsed > b.elapsed:
			return -1
		case a.elapsed < b.elapsed:
			return +1
		}
		return strings.Compare(a.name, b.name)
	})
	fmt.Printf("%d tests; slowest:\n", len(slowest))
	for _, r := range slowest[:min(5, len(slowest))] {
		fmt.Printf("  %6.3fs  %s\n", r.elapsed, r.name)
	}
}
