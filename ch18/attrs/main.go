// Attrs demonstrates attributes, groups, and the LogValuer interface.
package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/klbrg/gopl2/ch18/notime"
)

// A Token is a credential that must never appear in the logs.
type Token string

func (Token) LogValue() slog.Value { return slog.StringValue("REDACTED") }

// A User is expanded into a group of its own.
type User struct {
	Name  string
	ID    int
	Token Token
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("id", u.ID),
		slog.Any("token", u.Token),
	)
}

func main() {
	text := slog.New(slog.NewTextHandler(os.Stdout, notime.Options(nil)))
	json := slog.New(slog.NewJSONHandler(os.Stdout, notime.Options(nil)))
	u := User{Name: "alice", ID: 7, Token: "s3cr3t"}

	for _, logger := range []*slog.Logger{text, json} {
		logger = logger.With(slog.String("service", "orders"))
		logger.Info("request",
			slog.String("method", "GET"),
			slog.Group("http", "status", 200,
				"bytes", 1024,
				"elapsed", 750*time.Microsecond),
			slog.Any("user", u))

		conn := logger.WithGroup("conn")
		conn.LogAttrs(context.Background(), slog.LevelWarn, "reset",
			slog.String("peer", "10.0.0.2:443"),
			slog.Int("attempt", 3))

		logger.Info("batch", slog.GroupAttrs("stats",
			slog.Int("ok", 12), slog.Int("failed", 1)))
	}
}
