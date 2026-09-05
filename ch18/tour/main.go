// Tour prints one log record through each of the built-in handlers.
package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("open", "path", "/etc/passwd", "readonly", true)

	text := slog.New(slog.NewTextHandler(os.Stdout, nil))
	text.Info("open", "path", "/etc/passwd", "readonly", true)

	json := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	json.Info("open", "path", "/etc/passwd", "readonly", true)

	src := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{AddSource: true}))
	src.Info("open", "path", "/etc/passwd", "readonly", true)
}
