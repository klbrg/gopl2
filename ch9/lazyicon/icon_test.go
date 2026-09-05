package lazyicon

import (
	"sync"
	"testing"
)

func TestConcurrentIcon(t *testing.T) {
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if Icon("spades.png") == nil {
				t.Error("missing icon")
			}
		}()
	}
	wg.Wait()
}
