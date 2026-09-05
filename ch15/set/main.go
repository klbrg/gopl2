// Set implements a set of distinct elements of any comparable type.
package main

import (
	"fmt"
	"slices"
)

// !+set
// A Set is an unordered collection of distinct elements of type E.
// The zero Set is not usable; call NewSet.
type Set[E comparable] struct {
	m map[E]bool
}

func NewSet[E comparable](elems ...E) *Set[E] {
	s := &Set[E]{m: make(map[E]bool)}
	s.Add(elems...)
	return s
}

func (s *Set[E]) Add(elems ...E) {
	for _, e := range elems {
		s.m[e] = true
	}
}

func (s *Set[E]) Has(e E) bool { return s.m[e] }

func (s *Set[E]) Len() int { return len(s.m) }

// Elems returns the elements of s in unspecified order.
func (s *Set[E]) Elems() []E {
	elems := make([]E, 0, len(s.m))
	for e := range s.m {
		elems = append(elems, e)
	}
	return elems
}

// Union returns a new set containing every element of a or b.
func Union[E comparable](a, b *Set[E]) *Set[E] {
	u := NewSet(a.Elems()...)
	u.Add(b.Elems()...)
	return u
}

func main() {
	fruit := NewSet("apple", "pear", "apple")
	fmt.Println(fruit.Len(), fruit.Has("pear"), fruit.Has("fig")) // "2 true false"

	more := NewSet("fig")
	all := Union(fruit, more)
	elems := all.Elems()
	slices.Sort(elems)
	fmt.Println(elems) // "[apple fig pear]"
}
