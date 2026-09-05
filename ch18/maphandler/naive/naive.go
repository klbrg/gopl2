// Package naive contains a first, incorrect attempt at a slog.Handler
// that records log records in memory.  See the parent package for the
// version that passes testing/slogtest.
package naive

import (
	"context"
	"log/slog"
	"sync"
)

// A Handler records each log record as a map.
type Handler struct {
	mu      *sync.Mutex
	records *[]map[string]any
	attrs   []slog.Attr
	group   string
}

func New() *Handler {
	return &Handler{mu: new(sync.Mutex), records: new([]map[string]any)}
}

func (h *Handler) Records() []map[string]any {
	h.mu.Lock()
	defer h.mu.Unlock()
	return *h.records
}

func (h *Handler) Enabled(context.Context, slog.Level) bool { return true }

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := *h
	h2.attrs = append(h.attrs[:len(h.attrs):len(h.attrs)], attrs...)
	return &h2
}

func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := *h
	h2.group = name
	return &h2
}

func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	m := map[string]any{
		slog.TimeKey:    r.Time,
		slog.LevelKey:   r.Level,
		slog.MessageKey: r.Message,
	}
	add := func(a slog.Attr) {
		key := a.Key
		if h.group != "" {
			key = h.group + "." + key
		}
		m[key] = a.Value.Any()
	}
	for _, a := range h.attrs {
		add(a)
	}
	r.Attrs(func(a slog.Attr) bool { add(a); return true })

	h.mu.Lock()
	defer h.mu.Unlock()
	*h.records = append(*h.records, m)
	return nil
}
