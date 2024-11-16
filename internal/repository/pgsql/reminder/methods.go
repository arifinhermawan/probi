package reminder

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreateReminderInDB(ctx context.Context, req CreateReminderReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Database+"CreateReminderInDB")
	defer span.End()

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Timeout.FiveSeconds)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"user_id":   req.UserID,
		"title":     req.Title,
		"frequency": req.Frequency,
		"interval":  req.Interval,
	}

	namedQuery, args, err := sqlx.Named(queryCreateReminderInDB, req)
	if err != nil {
		log.Error(ctx, metadata, err, "[CreateReminderInDB] sqlx.Named() got error")
		return err
	}

	_, err = r.db.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Error(ctx, metadata, err, "[CreateReminderInDB] r.db.ExecContext() got error")
		return err
	}

	return nil
}

func (r *Repository) GetActiveReminderByUserIDFromDB(ctx context.Context, userID int64) ([]Reminder, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Database+"GetActiveReminderByUserIDFromDB")
	defer span.End()

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Timeout.FiveSeconds)*time.Second)
	defer cancel()

	var reminders []Reminder
	err := r.db.SelectContext(ctxTimeout, &reminders, queryGetActiveReminderByUserIDFromDB, userID)
	if err != nil && err != sql.ErrNoRows {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[GetActiveReminderByUserIDFromDB] r.db.SelectContext() got error")

		return nil, err
	}

	return reminders, nil
}

func (r *Repository) UpdateReminderInDB(ctx context.Context, req UpdateReminderReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Database+"UpdateReminderInDB")
	defer span.End()

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Timeout.FiveSeconds)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"id":        req.ID,
		"frequency": req.Frequency,
		"interval":  req.Interval,
		"end_date":  req.EndDate,
	}

	namedQuery, args, err := sqlx.Named(queryUpdateReminderInDB, req)
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateReminderInDB] sqlx.Named() got error")
		return err
	}

	_, err = r.db.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateReminderInDB] r.db.ExecContext() got error")
		return err
	}

	return nil
}
