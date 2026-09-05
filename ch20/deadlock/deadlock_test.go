// Package deadlock demonstrates the deadlock report of Section 20.6.
// The test panics by design, so it is behind a build tag:
//
//	go test -tags=deadlock gopl.io/ch20/deadlock
//
//go:build deadlock

package deadlock

import (
	"testing"
	"testing/synctest"
)

func TestDeadlock(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int)
		<-ch
	})
}
