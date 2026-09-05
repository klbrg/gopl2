#!/bin/sh
# Build and run every chapter 14 example.  Requires network on first run.
set -e
here=$(cd "$(dirname "$0")" && pwd)
export GOTOOLCHAIN=local

for m in hello twoversions whoami modinfo mvs/app; do
    echo "=== $m ==="
    (cd "$here/$m" && GOWORK=off go vet ./... && GOWORK=off go run .)
done

echo "=== workspace ==="
(cd "$here/workspace/hello" && go vet ./... && go run .)
