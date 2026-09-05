// Tree is an ordered binary tree whose in-order walk is an iter.Seq2.
package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

type Tree[K cmp.Ordered, V any] struct {
	left, right *Tree[K, V]
	key         K
	value       V
}

// Insert returns the tree t with key mapped to value.
func (t *Tree[K, V]) Insert(key K, value V) *Tree[K, V] {
	if t == nil {
		return &Tree[K, V]{key: key, value: value}
	}
	switch c := cmp.Compare(key, t.key); {
	case c < 0:
		t.left = t.left.Insert(key, value)
	case c > 0:
		t.right = t.right.Insert(key, value)
	default:
		t.value = value
	}
	return t
}

// walk reports whether the traversal should continue.
func (t *Tree[K, V]) walk(yield func(K, V) bool) bool {
	return t == nil ||
		t.left.walk(yield) && yield(t.key, t.value) && t.right.walk(yield)
}

// All returns an iterator over the tree's entries in increasing key order.
func (t *Tree[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { t.walk(yield) }
}

func main() {
	var t *Tree[string, int]
	for i, w := range []string{"pear", "fig", "quince", "apple", "medlar"} {
		t = t.Insert(w, i)
	}
	for k, v := range t.All() {
		fmt.Printf("%s=%d ", k, v)
	}
	fmt.Println()

	// break stops the recursion partway through.
	for k := range t.All() {
		fmt.Print(k, " ")
		if k == "medlar" {
			break
		}
	}
	fmt.Println()

	fmt.Println(slices.Sorted(func(yield func(string) bool) {
		for k := range t.All() {
			if !yield(k) {
				return
			}
		}
	}))
}
