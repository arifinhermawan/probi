package scheduler

import (
	"context"

	"github.com/arifinhermawan/probi/internal/app/cron/server"
)

func processDueReminder(ctx context.Context, handlers *server.Handlers) func() {
	return func() {
		handlers.Reminder.ProcessDueReminder(ctx)
	}
}
