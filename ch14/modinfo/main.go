// Modinfo prints the module dependencies recorded in each named executable.
package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
)

func main() {
	for _, name := range os.Args[1:] {
		info, err := buildinfo.ReadFile(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "modinfo: %v\n", err)
			continue
		}
		fmt.Printf("%s: %s %s\n", name, info.Path, info.Main.Version)
		for _, dep := range info.Deps {
			marker := " "
			if dep.Replace != nil {
				marker = "*"
			}
			fmt.Printf("  %s %-30s %s\n", marker, dep.Path, dep.Version)
		}
	}
}
