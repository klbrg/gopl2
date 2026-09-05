// Astype discriminates errors by type, the 1.13 way and the 1.26 way.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// A QueryError describes a failed query against a named table.
type QueryError struct {
	Table string
	Line  int
	Err   error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("query %s line %d: %v", e.Table, e.Line, e.Err)
}

func (e *QueryError) Unwrap() error { return e.Err }

// ErrSyntax reports a malformed query.
var ErrSyntax = errors.New("syntax error")

func query(table string) error {
	return fmt.Errorf("report %q: %w", "daily",
		&QueryError{Table: table, Line: 7, Err: ErrSyntax})
}

func main() {
	err := query("orders")
	fmt.Println(err)

	// Go 1.13: errors.As, with an out-parameter.
	var qe *QueryError
	if errors.As(err, &qe) {
		fmt.Println("table:", qe.Table, "line:", qe.Line)
	}

	// Go 1.26: errors.AsType, generic and type-safe.
	if qe, ok := errors.AsType[*QueryError](err); ok {
		fmt.Println("table:", qe.Table, "line:", qe.Line)
	}

	// The sentinel at the bottom of the tree is still reachable.
	fmt.Println(errors.Is(err, ErrSyntax))

	// The same works for the standard library's own types.
	_, ferr := os.Open("/no/such/file")
	if pe, ok := errors.AsType[*fs.PathError](ferr); ok {
		fmt.Printf("op=%s path=%s\n", pe.Op, pe.Path)
	}
}
