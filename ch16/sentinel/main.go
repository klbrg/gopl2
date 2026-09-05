// Sentinel shows how errors.Is searches a wrapped error's tree.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// ErrNoStore reports that no store has been configured.
var ErrNoStore = errors.New("no store configured")

type Store struct{ dir string }

// Load returns the contents of the named record.
func (s *Store) Load(name string) ([]byte, error) {
	if s.dir == "" {
		return nil, fmt.Errorf("load %s: %w", name, ErrNoStore)
	}
	data, err := os.ReadFile(s.dir + "/" + name)
	if err != nil {
		return nil, fmt.Errorf("load %s: %w", name, err)
	}
	return data, nil
}

func main() {
	var s Store
	_, err := s.Load("alice")
	fmt.Println(err)
	fmt.Println(err == ErrNoStore)              // equality fails
	fmt.Println(errors.Is(err, ErrNoStore))     // Is succeeds
	fmt.Println(errors.Is(err, fs.ErrNotExist)) // wrong sentinel

	s.dir = "/no/such/dir"
	_, err = s.Load("alice")
	fmt.Println(err)
	fmt.Println(errors.Is(err, fs.ErrNotExist))
	fmt.Println(errors.Is(err, ErrNoStore))
}
