package reminder

import (
	"context"
	"encoding/json"

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

	timeNow := svc.lib.GetTimeGMT7()
	request := reminder.CreateReminderReq{
		UserID:    req.UserID,
		Title:     req.Title,
		Frequency: req.Frequency,
		Interval:  req.Interval,
		StartDate: req.StartDate,
		DueDate:   dueDate,
		EndDate:   req.EndDate,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
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

func (svc *Service) ProcessDueReminder(ctx context.Context) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"ProcessDueReminder")
	defer span.End()

	res, err := svc.db.GetDueReminderIDsFromDB(ctx)
	if err != nil {
		log.Error(ctx, nil, err, "[ProcessDueReminder] svc.db.GetDueReminderIDsFromDB() got error")
		return err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Error(ctx, nil, err, "[ProcessDueReminder] json.Marshal() got error")
		return err
	}

	err = svc.nsq.PublishMessageToNSQ(ctx, svc.lib.GetConfig().Channel.Reminder.Topic, bytes)
	if err != nil {
		log.Error(ctx, nil, err, "[ProcessDueReminder] svc.nsq.PublishMessageToNSQ() got error")
		return err
	}

	return nil
}

func (svc *Service) GetUserActiveReminder(ctx context.Context, userID int64) ([]Reminder, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"GetUserActiveReminder")
	defer span.End()

	metadata := map[string]interface{}{
		"user_id": userID,
	}

	cached, err := svc.getReminderListFromRedis(ctx, userID)
	if err != nil {
		log.Warn(ctx, metadata, err, "[GetUserActiveReminder] svc.getReminderListFromRedis() got error")
	}

	if len(cached) != 0 {
		return cached, nil
	}

	res, err := svc.db.GetActiveReminderByUserIDFromDB(ctx, userID)
	if err != nil {
		log.Error(ctx, metadata, err, "[GetUserActiveReminder] svc.db.GetActiveReminderByUserIDFromDB() got error")
		return nil, err
	}

	reminders := make([]Reminder, len(res))
	for idx, reminder := range res {
		reminders[idx] = Reminder(reminder)
	}

	go func() {
		err := svc.setReminderListToRedis(ctx, userID, reminders)
		if err != nil {
			log.Warn(ctx, metadata, err, "[GetUserActiveReminder] svc.setReminderListToRedis() got error")
		}
	}()

	return reminders, nil
}

func (svc *Service) UpdateReminder(ctx context.Context, req UpdateReminderReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"UpdateReminder")
	defer span.End()

	timeNow := svc.lib.GetTimeGMT7()
	dueDate := svc.calculateDueDate(calculateDueDateReq{
		frequency: req.Frequency,
		interval:  req.Interval,
		startDate: timeNow,
		endDate:   req.EndDate,
	})

	request := reminder.UpdateReminderReq{
		ID:        req.ID,
		Frequency: req.Frequency,
		Interval:  req.Interval,
		DueDate:   dueDate,
		EndDate:   req.EndDate,
		UpdatedAt: timeNow,
	}

	metadata := map[string]interface{}{
		"id":        req.ID,
		"frequency": req.Frequency,
		"interval":  req.Interval,
		"due_date":  dueDate,
		"end_date":  req.EndDate,
	}

	err := svc.db.UpdateReminderInDB(ctx, request)
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateReminder] svc.db.UpdateReminderInDB() got error")
		return err
	}

	err = svc.deleteReminderListInRedis(ctx, req.UserID)
	if err != nil {
		log.Warn(ctx, metadata, err, "[UpdateReminder] svc.deleteReminderListInRedis() got error")
	}

	return nil
}
