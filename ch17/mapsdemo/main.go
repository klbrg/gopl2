// Mapsdemo exercises the maps package.
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	ages := map[string]int{"alice": 31, "bob": 20, "charlie": 34}

	// Keys and Values are iterators, not slices.
	fmt.Printf("%T\n", maps.Keys(ages))
	fmt.Println(slices.Sorted(maps.Keys(ages)))
	fmt.Println(slices.Sorted(maps.Values(ages)))

	// All yields pairs; sort the keys to get a deterministic order.
	for _, k := range slices.Sorted(maps.Keys(ages)) {
		fmt.Printf("%s:%d ", k, ages[k])
	}
	fmt.Println()

	// Collect builds a map from a Seq2; Insert adds to an existing one.
	swapped := maps.Collect(func(yield func(int, string) bool) {
		for k, v := range maps.All(ages) {
			if !yield(v, k) {
				return
			}
		}
	})
	fmt.Println(slices.Sorted(maps.Keys(swapped)))

	older := maps.Clone(ages)
	maps.Insert(older, maps.All(map[string]int{"bob": 21, "dora": 45}))
	fmt.Println(len(older), older["bob"], older["dora"])

	fmt.Println(maps.Equal(ages, maps.Clone(ages)))

	maps.DeleteFunc(older, func(k string, v int) bool { return v < 30 })
	fmt.Println(slices.Sorted(maps.Keys(older)))

	// Copy overwrites in place; Clone is nil-preserving.
	dst := map[string]int{"zoe": 1}
	maps.Copy(dst, ages)
	fmt.Println(len(dst))
	var nilmap map[string]int
	fmt.Println(maps.Clone(nilmap) == nil)

	clear(ages)
	fmt.Println(len(ages))
}
