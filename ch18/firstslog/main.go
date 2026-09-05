// Firstslog reports the progress of a batch of jobs using log/slog.
package main

import (
	"errors"
	"log/slog"
	"os"
	"time"
)

type job struct {
	id    int
	user  string
	bytes int
	err   error
}

var jobs = []job{
	{id: 1, user: "alice", bytes: 1024},
	{id: 2, user: "bob", err: errors.New("connection reset by peer")},
	{id: 3, user: "carol dean", bytes: 33},
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	for _, j := range jobs {
		start := time.Now()
		logger := logger.With("job", j.id, "user", j.user)
		if j.err != nil {
			logger.Error("job failed", "err", j.err)
			continue
		}
		logger.Info("job complete", "bytes", j.bytes,
			"elapsed", time.Since(start).Round(time.Millisecond))
	}
}
