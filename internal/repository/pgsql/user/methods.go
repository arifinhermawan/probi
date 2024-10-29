package user

import (
	"context"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreateUserInDB(ctx context.Context, req CreateUserReq) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"display_name": req.DisplayName,
		"email":        req.Email,
		"username":     req.Username,
	}

	namedQuery, args, err := sqlx.Named(queryCreateUserInDB, req)
	if err != nil {
		log.Error(ctx, metadata, err, "[CreateUserInDB] sqlx.Named() got error")
		return err
	}

	_, err = r.db.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Error(ctx, metadata, err, "[CreateUserInDB] r.db.ExecContext() got error")
		return err
	}

	return nil
}
