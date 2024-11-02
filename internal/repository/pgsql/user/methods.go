package user

import (
	"context"
	"database/sql"
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

func (r *Repository) GetUserByEmailFromDB(ctx context.Context, email string) (User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	var user User
	err := r.db.GetContext(ctxTimeout, &user, queryGetUserByEmailFromDB, email)
	if err != nil && err != sql.ErrNoRows {
		log.Error(ctx, map[string]interface{}{
			"email": email,
		}, err, "[GetUserByEmailFromDB] r.db.GetContext() got error")

		return User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByIDFromDB(ctx context.Context, userID int64) (User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	var user User
	err := r.db.GetContext(ctxTimeout, &user, queryGetUserByIDFromDB, userID)
	if err != nil && err != sql.ErrNoRows {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[GetUserByIDFromDB] r.db.GetContext() got error")

		return User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsernameFromDB(ctx context.Context, username string) (User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	var user User
	err := r.db.GetContext(ctxTimeout, &user, queryGetUserByUsernameFromDB, username)
	if err != nil && err != sql.ErrNoRows {
		log.Error(ctx, map[string]interface{}{
			"username": username,
		}, err, "[GetUserByUsernameFromDB] r.db.GetContext() got error")

		return User{}, err
	}

	return user, nil
}

func (r *Repository) UpdateUserDetailsInDB(ctx context.Context, tx *sql.Tx, req UpdateUserDetailsReq) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	namedQuery, args, err := sqlx.Named(queryUpdateUserInDB, req)
	if err != nil {
		log.Error(ctx, nil, err, "[UpdateUserDetailsInDB] sqlx.Named() got error")
		return err
	}

	_, err = tx.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Error(ctx, nil, err, "[UpdateUserDetailsInDB] tx.ExecContext() got error")
		return err
	}

	return nil
}
