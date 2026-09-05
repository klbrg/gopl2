// Package sum adds the elements of a slice.
package sum

// Sum returns the sum of the elements of xs.
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
