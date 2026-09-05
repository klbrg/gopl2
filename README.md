# The Go Programming Language, 2nd edition — example programs

Companion code for the second edition. Targets **Go 1.26**.

    go install github.com/klbrg/gopl2/ch1/helloworld@latest
    $(go env GOPATH)/bin/helloworld
    Hello, 世界

Or, working in the tree:

    git clone https://github.com/klbrg/gopl2
    cd gopl2
    go build ./...
    go run ./ch1/helloworld

## Provenance and licence

Chapters 1–13 derive from the first edition's companion programs at
<https://github.com/adonovan/gopl.io>, © 2016 Alan A. A. Donovan &
Brian W. Kernighan, licensed **CC BY-NC-SA 4.0**. This repository is a
derivative work and carries the same licence: non-commercial, share-alike,
attribution required. Chapters 14–20 are new.

## What differs from the first edition's code

- `io/ioutil` replaced throughout by `io` and `os` (deprecated in Go 1.16):
  `ch1/fetch`, `ch1/fetchall`, `ch1/dup3`, `ch8/du1`, `ch8/du2`, `ch8/du3`,
  `ch9/memotest`.
- `ch13/bzip` and `ch13/bzip-print` define their cgo helpers `static`, inside
  the preamble. In the first edition both defined a non-static `bz2compress` in
  a separate `.c` file; because `bzip-print`'s test imports `bzip`, both objects
  land in one binary and the link fails. `go build ./...` never catches this —
  only `go test ./...` does.
- `ch8/cake` calls `testing.Verbose()` from `TestMain` rather than during
  package initialization, which modern Go panics on.
- `ch12/methods`' example functions are named so `go vet` accepts them, and
  their expected output matches the five methods `time.Duration` has gained.
- The first edition's `//!+` / `//!-` excerpt markers are gone; the whole tree
  is `gofmt`-clean.
- Chapters are one module. `ch14` is the exception: minimal version selection
  and workspaces can only be demonstrated with several modules, so it has them.

## A note on `ch11/word1`

`go test ./...` reports one failure, in `ch11/word1`. That is deliberate and the
book depends on it: §11.2 shows those tests failing on "été" and on "A man, a
plan, a canal: Panama". `ch11/word2` is the corrected version.
