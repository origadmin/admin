package interrupt

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func WithContext(ctx context.Context) context.Context {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-signals
		cancel()
	}()
	return ctx
}
