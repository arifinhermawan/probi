package scheduler

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/cron/server"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/robfig/cron/v3"
)

func RegisterScheduler(ctx context.Context, schedule configuration.CronConfig, handlers *server.Handlers) {
	c := cron.New()

	_, err := c.AddFunc("* * * * *", processDueReminder(ctx, handlers))
	if err != nil {
		log.Error(ctx, nil, err, "[RegisterScheduler] failed to add ProcessDueReminder")
		return
	}

	c.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Info(ctx, nil, nil, "[RegisterScheduler] Shutdown signal received")

	c.Stop()
}
