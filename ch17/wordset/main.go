// Wordset defines a set type whose contents are exposed as an iter.Seq.
package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// A Set is an unordered collection of distinct comparable values.
type Set[E comparable] struct {
	m map[E]struct{}
}

func NewSet[E comparable](vals ...E) *Set[E] {
	s := &Set[E]{m: make(map[E]struct{})}
	for _, v := range vals {
		s.Add(v)
	}
	return s
}

func (s *Set[E]) Add(v E)           { s.m[v] = struct{}{} }
func (s *Set[E]) Contains(v E) bool { _, ok := s.m[v]; return ok }
func (s *Set[E]) Len() int          { return len(s.m) }

// All returns an iterator over the elements of s in no particular order.
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

// AddSeq adds every value of seq to s.
func (s *Set[E]) AddSeq(seq iter.Seq[E]) {
	for v := range seq {
		s.Add(v)
	}
}

func main() {
	s := NewSet("gopher", "walrus", "penguin")
	for v := range s.All() {
		_ = v // order is unspecified, so we don't print it here
	}
	fmt.Println(slices.Sorted(s.All()))

	t := NewSet[string]()
	t.AddSeq(s.All())
	t.AddSeq(slices.Values([]string{"otter", "walrus"}))
	fmt.Println(slices.Sorted(t.All()), t.Len())

	// A Set is built from a map, so maps.Keys is an iterator too.
	fmt.Println(slices.Sorted(maps.Keys(s.m)))
}
