package reminder

import (
	"context"
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

	timeNow := r.lib.GetTimeGMT7()
	param := map[string]interface{}{
		"user_id":    req.UserID,
		"title":      req.Title,
		"frequency":  req.Frequency,
		"interval":   req.Interval,
		"start_date": req.StartDate,
		"end_date":   req.EndDate,
		"due_date":   req.DueDate,
		"created_at": timeNow,
		"updated_at": timeNow,
	}

	namedQuery, args, err := sqlx.Named(queryCreateReminderInDB, param)
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
