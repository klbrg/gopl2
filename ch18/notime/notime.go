// Package notime provides slog handler options that suppress the time
// attribute, so that a program's log output is reproducible.
package notime

import "log/slog"

// Options returns handler options that log at level lvl and discard the
// top-level time attribute.  A nil lvl means slog.LevelInfo.
func Options(lvl slog.Leveler) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.TimeKey {
				return slog.Attr{} // discard
			}
			return a
		},
	}
}
