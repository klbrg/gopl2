// Slicesdemo exercises the slices package.
package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"delilah", "carol", "bob", "alice"}
	slices.Sort(names)
	fmt.Println(names)
	fmt.Println(slices.IsSorted(names), slices.Contains(names, "bob"))
	fmt.Println(slices.Index(names, "carol"), slices.Index(names, "zoe"))

	i, found := slices.BinarySearch(names, "carol")
	fmt.Println(i, found)
	i, found = slices.BinarySearch(names, "carla")
	fmt.Println(i, found)

	// Sorting by a key: cmp returns negative, zero, or positive.
	byLength := func(a, b string) int { return len(a) - len(b) }
	slices.SortStableFunc(names, byLength)
	fmt.Println(names)

	// Insert and Delete return the modified slice; the argument is spent.
	s := []int{1, 2, 5, 6}
	s = slices.Insert(s, 2, 3, 4)
	fmt.Println(s)
	s = slices.Delete(s, 1, 3)
	fmt.Println(s)
	s = slices.DeleteFunc(s, func(x int) bool { return x%2 == 0 })
	fmt.Println(s)

	// The aliasing hazard: Delete writes through the original array.
	orig := []string{"a", "b", "c", "d"}
	short := slices.Delete(orig, 1, 3)
	fmt.Printf("%q %q\n", short, orig)

	fmt.Println(slices.Equal([]int{1, 2}, []int{1, 2}))
	fmt.Println(slices.Compare([]int{1, 3}, []int{1, 2, 9}))
	fmt.Println(slices.Compact([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(slices.Concat([]int{1}, []int{2, 3}))
	fmt.Println(slices.Max([]int{4, 9, 2}), slices.Min([]int{4, 9, 2}))
	fmt.Println(slices.Repeat([]string{"ab"}, 3))

	slices.Reverse(names)
	fmt.Println(names)

	// Iterator-valued functions, all added in Go 1.23.
	for i, v := range slices.Backward([]string{"x", "y", "z"}) {
		fmt.Print(i, v, " ")
	}
	fmt.Println()
	fmt.Println(slices.Collect(slices.Values([]int{3, 1, 2})))
	fmt.Println(slices.Sorted(slices.Values([]int{3, 1, 2})))
	fmt.Println(slices.AppendSeq([]int{0}, slices.Values([]int{7, 8})))
	for chunk := range slices.Chunk([]int{1, 2, 3, 4, 5}, 2) {
		fmt.Print(chunk, " ")
	}
	fmt.Println()

	// A named slice type satisfies ~[]E, so it works unchanged.
	type Names []string
	var n Names = strings.Fields("gamma alpha beta")
	slices.Sort(n)
	fmt.Printf("%T %v\n", n, n)
}
