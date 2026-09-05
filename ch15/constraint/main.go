// Constraint demonstrates constraints that specify methods, including
// the pointer-receiver idiom.
package main

import (
	"fmt"
	"strconv"
)

// A Named is any type with a Name method.
type Named interface {
	Name() string
}

// Names returns the names of the elements of s.
// Because s is a []T and not a []Named, no boxing takes place.
func Names[T Named](s []T) []string {
	names := make([]string, len(s))
	for i, v := range s {
		names[i] = v.Name()
	}
	return names
}

// A Setter[T] is a pointer to T whose method Set parses a string.
// The *T element restricts the type set to pointers to T; the Set
// method makes the method callable on a value of type P.
type Setter[T any] interface {
	*T
	Set(string) error
}

// ParseAll parses each input into a fresh T. P must be given
// explicitly; T is then inferred from P's constraint.
func ParseAll[P Setter[T], T any](inputs []string) ([]T, error) {
	out := make([]T, len(inputs))
	for i, s := range inputs {
		if err := P(&out[i]).Set(s); err != nil {
			return nil, err
		}
	}
	return out, nil
}

type Employee struct{ First, Last string }

func (e Employee) Name() string { return e.First + " " + e.Last }

type Celsius float64

func (c *Celsius) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*c = Celsius(f)
	return nil
}

func main() {
	staff := []Employee{{"Ada", "Lovelace"}, {"Alan", "Turing"}}
	fmt.Println(Names(staff)) // "[Ada Lovelace Alan Turing]"

	temps, err := ParseAll[*Celsius]([]string{"-40", "37.5"})
	fmt.Println(temps, err) // "[-40 37.5] <nil>"

	_, err = ParseAll[*Celsius]([]string{"warm"})
	fmt.Println(err) // "strconv.ParseFloat: parsing \"warm\": invalid syntax"
}
