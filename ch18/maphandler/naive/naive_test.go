package naive

import (
	"testing"
	"testing/slogtest"
)

// TestNaive records the ways in which this handler is wrong.
func TestNaive(t *testing.T) {
	h := New()
	err := slogtest.TestHandler(h, h.Records)
	if err == nil {
		t.Fatal("naive handler unexpectedly passed slogtest")
	}
	t.Log("\n" + err.Error())
}
