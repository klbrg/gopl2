// Alias demonstrates generic type aliases, fully supported since Go 1.24.
package main

import (
	"cmp"
	"fmt"
	"slices"
)

// !+decls
// A Pair is a key and its associated value.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Pairs is a generic alias: it must be instantiated where it is used,
// and Pairs[K, V] is identical to []Pair[K, V].
type Pairs[K comparable, V any] = []Pair[K, V]

// Counts is a generic alias for a map with string keys.
type Counts[V any] = map[string]V

// WordCounts is an ordinary alias for one instantiation.
type WordCounts = Counts[int]

// !+sort
// SortByKey sorts ps by key. Its parameter is written []Pair[K, V],
// yet a Pairs[K, V] may be passed without conversion: the alias
// denotes the very same type.
func SortByKey[K cmp.Ordered, V any](ps []Pair[K, V]) {
	slices.SortFunc(ps, func(a, b Pair[K, V]) int {
		return cmp.Compare(a.Key, b.Key)
	})
}

func main() {
	wc := WordCounts{"the": 12, "a": 5}
	fmt.Println(wc["the"], len(wc)) // "12 2"

	ps := Pairs[string, int]{{"pear", 2}, {"apple", 1}}
	SortByKey(ps)
	fmt.Println(ps) // "[{apple 1} {pear 2}]"
}
