// Whoami reports how the running executable was built.
package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "whoami: no build information")
		os.Exit(1)
	}
	fmt.Printf("toolchain %s\n", info.GoVersion)
	fmt.Printf("main      %s %s\n", info.Main.Path, info.Main.Version)
	for _, dep := range info.Deps {
		fmt.Printf("dep       %s %s\n", dep.Path, dep.Version)
	}
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs", "vcs.revision", "vcs.time", "vcs.modified", "-trimpath":
			fmt.Printf("setting   %s=%s\n", s.Key, s.Value)
		}
	}
}
