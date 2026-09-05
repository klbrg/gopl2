// Package config reads a simple key=value settings file.
package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Path reports the settings file to read: the value of GOPLCONF
// if it is set, and otherwise gopl.conf in the current directory.
func Path() string {
	if p := os.Getenv("GOPLCONF"); p != "" {
		return p
	}
	return "gopl.conf"
}

// Load reads the settings file named by Path.
// Blank lines and lines beginning with # are ignored.
func Load() (map[string]string, error) {
	f, err := os.Open(Path())
	if err != nil {
		return nil, err
	}
	defer f.Close()
	settings := make(map[string]string)
	in := bufio.NewScanner(f)
	for line := 1; in.Scan(); line++ {
		text := strings.TrimSpace(in.Text())
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}
		key, value, ok := strings.Cut(text, "=")
		if !ok {
			return nil, fmt.Errorf("%s:%d: missing '='", Path(), line)
		}
		settings[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	return settings, in.Err()
}
