// Leak contrasts an error API that exposes its implementation with one that does not.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// ErrNotFound reports that no record exists for a key. It belongs to the
// Store API and does not change when the storage medium does.
var ErrNotFound = errors.New("store: record not found")

// A Store maps keys to records.
type Store interface {
	Get(key string) ([]byte, error)
}

// A fileStore keeps one record per file.
type fileStore struct{ dir string }

// naiveGet passes the file system's error straight through.
func (s *fileStore) naiveGet(key string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(s.dir, key))
	if err != nil {
		return nil, fmt.Errorf("get %s: %w", key, err)
	}
	return data, nil
}

// Get translates the file system's error into the Store's own vocabulary.
func (s *fileStore) Get(key string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(s.dir, key))
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return nil, fmt.Errorf("get %s: %w", key, ErrNotFound)
	case err != nil:
		return nil, fmt.Errorf("get %s: %v", key, err) // detail, not commitment
	}
	return data, nil
}

// A memStore keeps records in memory.
type memStore struct{ m map[string][]byte }

func (s *memStore) Get(key string) ([]byte, error) {
	data, ok := s.m[key]
	if !ok {
		return nil, fmt.Errorf("get %s: %w", key, ErrNotFound)
	}
	return data, nil
}

// lookup is written against the Store API alone.
func lookup(s Store, key string) string {
	_, err := s.Get(key)
	switch {
	case errors.Is(err, ErrNotFound):
		return "absent"
	case err != nil:
		return "failed: " + err.Error()
	}
	return "present"
}

func main() {
	dir, _ := os.MkdirTemp("", "ch16")
	defer os.RemoveAll(dir)
	os.WriteFile(filepath.Join(dir, "bob"), []byte("hi"), 0o666)

	fs1 := &fileStore{dir: dir}
	_, err := fs1.naiveGet("alice")
	fmt.Println(err)
	fmt.Println("naiveGet is fs.ErrNotExist:", errors.Is(err, fs.ErrNotExist))

	_, err = fs1.Get("alice")
	fmt.Println(err)
	fmt.Println("Get is fs.ErrNotExist:", errors.Is(err, fs.ErrNotExist))
	fmt.Println("Get is ErrNotFound:  ", errors.Is(err, ErrNotFound))

	mem := &memStore{m: map[string][]byte{"bob": []byte("hi")}}
	for _, s := range []Store{fs1, mem} {
		fmt.Printf("%-16T alice=%s bob=%s\n", s, lookup(s, "alice"), lookup(s, "bob"))
	}
}
