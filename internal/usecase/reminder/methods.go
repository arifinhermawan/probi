package reminder

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/service/reminder"
)

func (uc *UseCase) CreateReminder(ctx context.Context, req CreateReminderReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.UseCase+"CreateReminder")
	defer span.End()

	err := uc.reminder.CreateReminder(ctx, reminder.CreateReminderReq(req))
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id":   req.UserID,
			"title":     req.Title,
			"frequency": req.Frequency,
			"interval":  req.Interval,
		}, err, "[CreateReminder] uc.reminder.CreateReminder() got error")
		return err
	}

	return nil
}

func (uc *UseCase) GetUserActiveReminder(ctx context.Context, userID int64) ([]Reminder, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.UseCase+"GetUserActiveReminder")
	defer span.End()

	res, err := uc.reminder.GetUserActiveReminder(ctx, userID)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[GetUserActiveReminder] uc.reminder.GetUserActiveReminder() got error")

		return nil, err
	}

	reminders := make([]Reminder, len(res))
	for idx, reminder := range res {
		reminders[idx] = Reminder(reminder)
	}

	return reminders, nil
}
