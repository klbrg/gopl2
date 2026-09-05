// Package maphandler provides a slog.Handler that records each log
// record in memory as a map, for use in tests.
package maphandler

import (
	"context"
	"log/slog"
	"sync"
)

// A Handler implements slog.Handler by appending each record to a
// shared slice of maps.  Groups become nested maps.
type Handler struct {
	level   slog.Leveler
	mu      *sync.Mutex       // guards records; shared with derived handlers
	records *[]map[string]any // shared with derived handlers
	goas    []groupOrAttrs    // open groups and attributes, in order
}

// A groupOrAttrs holds either a group name or a list of attributes,
// recording the order in which WithGroup and WithAttrs were called.
type groupOrAttrs struct {
	group string      // group name if non-empty
	attrs []slog.Attr // attributes if non-empty
}

// New returns a Handler that records events at level lvl and above.
// A nil lvl means slog.LevelInfo.
func New(lvl slog.Leveler) *Handler {
	return &Handler{
		level:   lvl,
		mu:      new(sync.Mutex),
		records: new([]map[string]any),
	}
}

// Records returns the records logged so far.
func (h *Handler) Records() []map[string]any {
	h.mu.Lock()
	defer h.mu.Unlock()
	return *h.records
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	min := slog.LevelInfo
	if h.level != nil {
		min = h.level.Level()
	}
	return level >= min
}

func (h *Handler) withGroupOrAttrs(goa groupOrAttrs) *Handler {
	h2 := *h // copies level, mu, and records; the slice is copied below
	h2.goas = make([]groupOrAttrs, len(h.goas)+1)
	copy(h2.goas, h.goas)
	h2.goas[len(h2.goas)-1] = goa
	return &h2
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
}

func (h *Handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{group: name})
}

func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	rec := map[string]any{
		slog.LevelKey:   r.Level,
		slog.MessageKey: r.Message,
	}
	if !r.Time.IsZero() {
		rec[slog.TimeKey] = r.Time
	}
	if r.PC != 0 {
		if src := r.Source(); src != nil {
			rec[slog.SourceKey] = src
		}
	}

	// Walk the open groups, creating a nested map for each one.
	type open struct {
		parent map[string]any
		name   string
		child  map[string]any
	}
	var groups []open
	m := rec
	for _, goa := range h.goas {
		if goa.group != "" {
			child := map[string]any{}
			groups = append(groups, open{m, goa.group, child})
			m = child
		} else {
			for _, a := range goa.attrs {
				addAttr(m, a)
			}
		}
	}
	r.Attrs(func(a slog.Attr) bool { addAttr(m, a); return true })

	// Attach each group to its parent, innermost first, but only if
	// it has attributes.
	for i := len(groups) - 1; i >= 0; i-- {
		if g := groups[i]; len(g.child) > 0 {
			g.parent[g.name] = g.child
		}
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	*h.records = append(*h.records, rec)
	return nil
}

// addAttr adds a to m, resolving its value and flattening groups.
func addAttr(m map[string]any, a slog.Attr) {
	a.Value = a.Value.Resolve()
	if a.Equal(slog.Attr{}) {
		return // ignore empty Attr
	}
	if a.Value.Kind() != slog.KindGroup {
		m[a.Key] = a.Value.Any()
		return
	}
	attrs := a.Value.Group()
	if len(attrs) == 0 {
		return // ignore empty group
	}
	if a.Key == "" { // inline a group with no name
		for _, a := range attrs {
			addAttr(m, a)
		}
		return
	}
	child := map[string]any{}
	for _, a := range attrs {
		addAttr(child, a)
	}
	if len(child) > 0 {
		m[a.Key] = child
	}
}
