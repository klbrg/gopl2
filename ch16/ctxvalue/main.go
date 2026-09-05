// Ctxvalue carries a request ID across an API boundary.
package main

import (
	"context"
	"fmt"
)

// contextKey is unexported, so no other package can collide with it.
type contextKey int

const requestIDKey contextKey = 0

// WithRequestID returns a copy of ctx carrying the request ID id.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestID returns the request ID carried by ctx, if any.
func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

// logf prefixes each message with the request ID, when there is one.
func logf(ctx context.Context, format string, args ...any) {
	if id, ok := RequestID(ctx); ok {
		format = "[" + id + "] " + format
	}
	fmt.Printf(format+"\n", args...)
}

func handle(ctx context.Context, path string) {
	logf(ctx, "handling %s", path)
	store(ctx, path)
}

func store(ctx context.Context, key string) {
	logf(ctx, "writing %s", key)
}

func main() {
	handle(context.Background(), "/health")
	handle(WithRequestID(context.Background(), "a91f"), "/orders")

	// A key of a different type does not collide.
	type otherKey int
	ctx := context.WithValue(WithRequestID(context.Background(), "a91f"),
		otherKey(0), "shadow")
	id, ok := RequestID(ctx)
	fmt.Println(id, ok)
}
