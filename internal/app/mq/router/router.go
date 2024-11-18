package router

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arifinhermawan/probi/internal/app/mq/server"
)

func RegisterConsumer(ctx context.Context, h *server.Handlers) {
	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		h.Reminder.SendReminderEmail(ctx)
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	cancel()
	wg.Wait()
}
