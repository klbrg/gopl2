// Levels demonstrates dynamic control of the logging level.
package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

// LevelTrace is more verbose than any of slog's predefined levels.
const LevelTrace = slog.Level(-8)

var programLevel = new(slog.LevelVar) // LevelInfo by default

func replace(groups []string, a slog.Attr) slog.Attr {
	if len(groups) > 0 {
		return a
	}
	switch a.Key {
	case slog.TimeKey:
		return slog.Attr{} // discard, so that this output is reproducible
	case slog.LevelKey:
		if a.Value.Any().(slog.Level) == LevelTrace {
			a.Value = slog.StringValue("TRACE")
		}
	}
	return a
}

func main() {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       programLevel,
		ReplaceAttr: replace,
	})
	slog.SetDefault(slog.New(h))
	slog.SetLogLoggerLevel(slog.LevelWarn)

	ctx := context.Background()
	slog.Info("listening", "addr", ":8000")
	slog.Debug("cache miss", "key", "u/42") // discarded
	log.Print("started by the log package")

	programLevel.Set(LevelTrace)
	slog.Debug("cache miss", "key", "u/42")
	slog.Log(ctx, LevelTrace, "syscall", "name", "epoll_wait")

	if slog.Default().Enabled(ctx, LevelTrace) {
		slog.Log(ctx, LevelTrace, "table", "rows", census())
	}
}

func census() int { return 41 } // pretend this is expensive
