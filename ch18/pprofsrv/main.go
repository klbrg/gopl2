// Pprofsrv is a server whose profiles may be fetched over HTTP.
package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/pprof"
	"strconv"

	_ "net/http/pprof" // registers handlers under /debug/pprof/
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	http.HandleFunc("/hash", hash)
	slog.Info("listening", "addr", "localhost:6060")
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func hash(w http.ResponseWriter, r *http.Request) {
	tenant := r.FormValue("tenant")
	if tenant == "" {
		tenant = "anonymous"
	}
	n, err := strconv.Atoi(r.FormValue("n"))
	if err != nil {
		http.Error(w, "bad n", http.StatusBadRequest)
		return
	}
	labels := pprof.Labels("tenant", tenant, "endpoint", "/hash")
	pprof.Do(r.Context(), labels, func(ctx context.Context) {
		fmt.Fprintf(w, "%x\n", digest(n))
	})
}

// digest hashes a small buffer n times.
func digest(n int) [32]byte {
	sum := [32]byte{}
	for i := 0; i < n; i++ {
		sum = sha256.Sum256(sum[:])
	}
	return sum
}
