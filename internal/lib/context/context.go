package context

import (
	"context"
	"os"

	"github.com/arifinhermawan/blib/log"
)

func DefaultContext() context.Context {
	ctx := context.Background()

	hostname, _ := os.Hostname()
	service := "probi"

	ctx = context.WithValue(ctx, log.ContextKey("host"), hostname)
	ctx = context.WithValue(ctx, log.ContextKey("service"), service)

	return ctx
}
