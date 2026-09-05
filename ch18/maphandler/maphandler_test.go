package maphandler

import (
	"log/slog"
	"testing"
	"testing/slogtest"
)

func TestSlogtest(t *testing.T) {
	var h *Handler
	slogtest.Run(t,
		func(*testing.T) slog.Handler {
			h = New(nil)
			return h
		},
		func(*testing.T) map[string]any {
			return h.Records()[0]
		})
}
