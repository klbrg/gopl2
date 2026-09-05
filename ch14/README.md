# Chapter 14 examples

Verified with go1.26.7 on linux/amd64.  Every command below is run with
`GOWORK=off` except inside `workspace/`, which is the workspace example.

    hello/        14.1, 14.2, 14.5, 14.6  a module with one dependency
    twoversions/  14.4                    v1 and v3 of one library at once
    mvs/          14.3                    minimal version selection, offline
    workspace/    14.7                    a two-module go.work workspace
    whoami/       14.10                   runtime/debug.ReadBuildInfo
    modinfo/      14.10                   debug/buildinfo.ReadFile

Build everything:

    ./verify.sh
