// Wordfreq prints the most frequent words in its input as a bar chart.
package main

import (
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		w := strings.ToLower(strings.TrimFunc(in.Text(), func(r rune) bool {
			return !unicode.IsLetter(r)
		}))
		if w != "" {
			counts[w]++
		}
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "wordfreq:", err)
		os.Exit(1)
	}

	// Most frequent first; ties broken alphabetically.
	words := slices.SortedFunc(maps.Keys(counts), func(a, b string) int {
		return cmp.Or(
			cmp.Compare(counts[b], counts[a]),
			cmp.Compare(a, b),
		)
	})

	for _, w := range words[:min(10, len(words))] {
		fmt.Printf("%-8s %s\n", w, strings.Repeat("*", counts[w]))
	}
}
