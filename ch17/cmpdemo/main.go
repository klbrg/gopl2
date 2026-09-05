// Cmpdemo exercises the cmp package and the min, max, and clear builtins.
package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	fmt.Println(min(3, 1, 2), max(3, 1, 2))
	fmt.Println(min(2.5, 1.0), max("bee", "ant"))

	nan := math.NaN()
	fmt.Println(nan < 0.0, nan == nan)                        // both false
	fmt.Println(cmp.Compare(nan, 0.0), cmp.Compare(nan, nan)) // -1 0
	fmt.Println(cmp.Less(nan, 0.0), cmp.Less(1, 2))

	fmt.Println(cmp.Or("", "", "fallback"))
	fmt.Println(cmp.Or(0, 7, 9))
	fmt.Printf("%q\n", cmp.Or[string]())

	// The multi-tier ordering of Section 7.6, written with Or.
	slices.SortFunc(tracks, func(x, y *Track) int {
		return cmp.Or(
			cmp.Compare(x.Title, y.Title),
			cmp.Compare(x.Year, y.Year),
			cmp.Compare(x.Length, y.Length),
		)
	})
	printTracks(tracks)

	// clear empties a map, including keys that delete cannot name.
	m := map[float64]string{1: "one", nan: "nan"}
	delete(m, nan)
	fmt.Println(len(m))
	clear(m)
	fmt.Println(len(m), m == nil)

	// On a slice, clear zeroes the elements; the length is unchanged.
	s := []string{"a", "b", "c"}
	clear(s)
	fmt.Printf("%q %d\n", s, len(s))
}
