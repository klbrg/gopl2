package bank

import (
	"log/slog"
	"testing"

	"github.com/klbrg/gopl2/ch18/maphandler"
)

func TestWithdrawIsAudited(t *testing.T) {
	h := maphandler.New(slog.LevelInfo)
	a := &Account{Log: slog.New(h)}
	a.Deposit(100)
	if err := a.Withdraw(500); err != ErrInsufficient {
		t.Fatalf("Withdraw(500) = %v, want %v", err, ErrInsufficient)
	}

	records := h.Records()
	if len(records) != 2 {
		t.Fatalf("logged %d records, want 2", len(records))
	}
	got := records[1]
	if got[slog.MessageKey] != "withdrawal refused" {
		t.Errorf("msg = %q", got[slog.MessageKey])
	}
	if got[slog.LevelKey] != slog.LevelWarn {
		t.Errorf("level = %v, want WARN", got[slog.LevelKey])
	}
	// slog.Int stores an int64, so that is what comes back out.
	if got["balance"] != int64(100) {
		t.Errorf("balance = %[1]v (%[1]T), want 100", got["balance"])
	}
}
