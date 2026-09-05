// Wrap demonstrates the %w verb and the chain it builds.
package main

import (
	"errors"
	"fmt"
	"os"
)

// readConfig reads a configuration file.
func readConfig(path string) error {
	if _, err := os.ReadFile(path); err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	return nil
}

// startup prepares the server.
func startup(path string) error {
	if err := readConfig(path); err != nil {
		return fmt.Errorf("startup: %w", err)
	}
	return nil
}

func main() {
	err := startup("/no/such/file")
	fmt.Println(err)
	for e := err; e != nil; e = errors.Unwrap(e) {
		fmt.Printf("%-24T %v\n", e, e)
	}
}
