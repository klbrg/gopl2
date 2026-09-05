package bench

import (
	"context"
	"log"
	"log/slog"
	"testing"
	"time"
)

// A sink accepts and discards output.  We do not use io.Discard,
// because the log package tests for it and returns early.
type sink struct{}

func (sink) Write(p []byte) (int, error) { return len(p), nil }

var (
	ctx     = context.Background()
	out     = sink{}
	plain   = log.New(out, "", log.LstdFlags)
	text    = slog.New(slog.NewTextHandler(out, nil))
	json    = slog.New(slog.NewJSONHandler(out, nil))
	elapsed = 1500 * time.Microsecond
)

func BenchmarkLogPrintf(b *testing.B) {
	for b.Loop() {
		plain.Printf("request done method=%s status=%d elapsed=%v",
			"GET", 200, elapsed)
	}
}

func BenchmarkTextArgs(b *testing.B) {
	for b.Loop() {
		text.Info("request done", "method", "GET", "status", 200,
			"elapsed", elapsed)
	}
}

func BenchmarkTextAttrs(b *testing.B) {
	for b.Loop() {
		text.LogAttrs(ctx, slog.LevelInfo, "request done",
			slog.String("method", "GET"), slog.Int("status", 200),
			slog.Duration("elapsed", elapsed))
	}
}

func BenchmarkJSONAttrs(b *testing.B) {
	for b.Loop() {
		json.LogAttrs(ctx, slog.LevelInfo, "request done",
			slog.String("method", "GET"), slog.Int("status", 200),
			slog.Duration("elapsed", elapsed))
	}
}

func BenchmarkDisabled(b *testing.B) {
	for b.Loop() {
		text.Debug("request done", "method", "GET", "status", 200,
			"elapsed", elapsed)
	}
}

func BenchmarkDisabledAttrs(b *testing.B) {
	for b.Loop() {
		text.LogAttrs(ctx, slog.LevelDebug, "request done",
			slog.String("method", "GET"), slog.Int("status", 200),
			slog.Duration("elapsed", elapsed))
	}
}

// A logger derived with With formats its attributes only once.
var withLogger = text.With(slog.String("service", "orders"),
	slog.String("region", "eu-north-1"))

func BenchmarkWith(b *testing.B) {
	for b.Loop() {
		withLogger.LogAttrs(ctx, slog.LevelInfo, "request done",
			slog.String("method", "GET"))
	}
}

func BenchmarkWithoutWith(b *testing.B) {
	for b.Loop() {
		text.LogAttrs(ctx, slog.LevelInfo, "request done",
			slog.String("service", "orders"),
			slog.String("region", "eu-north-1"),
			slog.String("method", "GET"))
	}
}

// expensive is a value whose string form is costly to compute.
type expensive struct{ n int }

func (e expensive) String() string {
	s := ""
	for i := 0; i < e.n; i++ {
		s += "x"
	}
	return s
}

func (e expensive) LogValue() slog.Value { return slog.StringValue(e.String()) }

func BenchmarkDisabledEager(b *testing.B) {
	e := expensive{n: 64}
	for b.Loop() {
		text.Debug("dump", "value", e.String())
	}
}

func BenchmarkDisabledLazy(b *testing.B) {
	e := expensive{n: 64}
	for b.Loop() {
		text.LogAttrs(ctx, slog.LevelDebug, "dump", slog.Any("value", e))
	}
}
