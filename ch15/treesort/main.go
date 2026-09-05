// Treesort sorts a slice using an ordered binary tree.
// It is the generic counterpart of gopl.io/ch4/treesort.
package main

import (
	"cmp"
	"fmt"
)

// !+tree
type tree[T cmp.Ordered] struct {
	left, right *tree[T]
	value       T
}

// Sort sorts values in place.
func Sort[S ~[]E, E cmp.Ordered](values S) {
	var root *tree[E]
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues[T cmp.Ordered](values []T, t *tree[T]) []T {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add[T cmp.Ordered](t *tree[T], value T) *tree[T] {
	if t == nil {
		// Equivalent to return &tree[T]{value: value}.
		t = new(tree[T])
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// !+words
// A Words is a slice of strings, so ~[]E admits it.
type Words []string

func main() {
	nums := []int{5, 3, 8, 1, 9, 2}
	Sort(nums)
	fmt.Println(nums) // "[1 2 3 5 8 9]"

	w := Words{"delta", "alpha", "charlie", "bravo"}
	Sort(w)
	fmt.Println(w) // "[alpha bravo charlie delta]"
}
