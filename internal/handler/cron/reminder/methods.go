package reminder

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/lib/constants"
)

func (h *Handler) ProcessDueReminder(ctx context.Context) {
	ctx, txn := tracer.StartTransaction(ctx, tracer.Handler+"ProcessDueReminder", constants.ServiceNameCRON)
	defer txn.End()

	metadata := map[string]interface{}{
		"cron_name": "ProcessDueReminder",
	}

	log.Info(ctx, metadata, nil, "[ProcessDueReminder] CRON job started")
	err := h.reminder.ProcessDueReminder(ctx)
	if err != nil {
		log.Error(ctx, nil, err, "[ProcessDueReminder] h.reminder.ProcessDueReminder() got error")
		return
	}

	log.Info(ctx, metadata, nil, "[ProcessDueReminder] CRON finished")
}
