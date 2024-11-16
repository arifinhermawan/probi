package reminder

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/reminder"
)

func (svc *Service) CreateReminder(ctx context.Context, req CreateReminderReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"CreateReminder")
	defer span.End()

	dueDate := svc.calculateDueDate(calculateDueDateReq{
		frequency: req.Frequency,
		interval:  req.Interval,
		startDate: req.StartDate,
		endDate:   req.EndDate,
	})

	request := reminder.CreateReminderReq{
		UserID:    req.UserID,
		Title:     req.Title,
		Frequency: req.Frequency,
		Interval:  req.Interval,
		StartDate: req.StartDate,
		DueDate:   dueDate,
		EndDate:   req.EndDate,
	}

	metadata := map[string]interface{}{
		"user_id":   req.UserID,
		"title":     req.Title,
		"frequency": req.Frequency,
		"interval":  req.Interval,
	}

	err := svc.db.CreateReminderInDB(ctx, request)
	if err != nil {
		log.Error(ctx, metadata, err, "[CreateReminder] svc.db.CreateReminderInDB() got error")
		return err
	}

	err = svc.deleteReminderListInRedis(ctx, req.UserID)
	if err != nil {
		log.Warn(ctx, metadata, err, "[CreateReminder] svc.deleteReminderListInRedis() got error")
	}

	return nil
}
