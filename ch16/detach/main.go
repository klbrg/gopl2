// Detach shows context.AfterFunc and context.WithoutCancel.
package main

import (
	"context"
	"fmt"
	"time"
)

// audit writes a record that must survive cancellation of the request.
func audit(ctx context.Context, msg string) {
	// ctx has the request's values but not its cancellation.
	ctx = context.WithoutCancel(ctx)
	fmt.Println("audit:", msg, "err =", ctx.Err())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("cleanup running")
		close(done)
	})
	defer stop()

	cancel()
	<-done
	audit(ctx, "order 42 abandoned")

	// A context already done runs the function immediately.
	ran := make(chan struct{})
	context.AfterFunc(ctx, func() { close(ran) })
	select {
	case <-ran:
		fmt.Println("ran immediately")
	case <-time.After(time.Second):
		fmt.Println("timed out")
	}
}
