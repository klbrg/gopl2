// Ctxcancel shows a generator goroutine that stops when its context is done.
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// squares sends 1, 4, 9, ... on the returned channel until ctx is done.
func squares(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := 1; ; n++ {
			select {
			case ch <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for v := range squares(ctx) {
		fmt.Print(v, " ")
		if v > 50 {
			break
		}
	}
	fmt.Println()

	fmt.Println("before cancel:", ctx.Err(), runtime.NumGoroutine())
	cancel()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("after cancel: ", ctx.Err(), runtime.NumGoroutine())
	fmt.Println("cause:        ", context.Cause(ctx))
}
