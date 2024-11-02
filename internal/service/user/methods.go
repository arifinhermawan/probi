package user

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
)

func (svc *Service) CreateUser(ctx context.Context, req CreateUserReq) error {
	timeNow := svc.lib.GetTimeGMT7()

	err := svc.db.CreateUserInDB(ctx, user.CreateUserReq{
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Email:       req.Email,
		Password:    req.Password,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	})
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"display_name": req.DisplayName,
			"email":        req.Email,
			"username":     req.Username,
		}, err, "[CreateUser] svc.db.CreateUserInDB() got error")
		return err
	}

	return nil
}

func (svc *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := svc.db.GetUserByEmailFromDB(ctx, email)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"email": email,
		}, err, "[GetUserByEmail] svc.db.GetUserByEmailFromDB() got error")

		return User{}, err
	}

	return User(user), nil
}

func (svc *Service) GetUserByID(ctx context.Context, userID int64) (User, error) {
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	user, err := svc.getUserDetailFromRedis(ctx, userID)
	if err != nil {
		log.Warn(ctx, metadata, err, "[GetUserByID] svc.getUserDetailFromRedis() got error")
	}

	if user.ID != 0 {
		return user, nil
	}

	res, err := svc.db.GetUserByIDFromDB(ctx, userID)
	if err != nil {
		log.Error(ctx, metadata, err, "[GetUserByID] svc.db.GetUserByIDFromDB() got error")
		return User{}, err
	}

	details := User{
		ID:          res.ID,
		Username:    res.Username,
		DisplayName: res.DisplayName,
		Email:       res.Email,
	}

	go func() {
		err := svc.setUserDetailToRedis(ctx, details)
		if err != nil {
			log.Warn(ctx, metadata, err, "[GetUserByID] svc.setUserDetailToRedis() got error")
		}
	}()

	return details, nil
}

func (svc *Service) GetUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := svc.db.GetUserByUsernameFromDB(ctx, username)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"username": username,
		}, err, "[GetUserByEmail] svc.db.GetUserByUsernameFromDB() got error")

		return User{}, err
	}

	return User(user), nil
}

func (svc *Service) UpdateUserDetails(ctx context.Context, req UpdateUserDetailsReq) error {
	metadata := map[string]interface{}{
		"user_id": req.UserID,
	}

	var err error
	tx, err := svc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateUserDetails] svc.db.BeginTx() got error")
		return err
	}

	defer func() {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				log.Error(ctx, metadata, err, "[UpdateUserDetails] tx.Rollback() got error")
			}
		}
	}()

	err = svc.db.UpdateUserDetailsInDB(ctx, tx, user.UpdateUserDetailsReq{
		UserID:      req.UserID,
		Username:    req.Username,
		DisplayName: req.DisplayName,
		UpdatedAt:   svc.lib.GetTimeGMT7(),
	})
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateUserDetails] svc.db.UpdateUserDetailsInDB() got error")
		return err
	}

	details := User{
		ID:          req.UserID,
		Username:    req.Username,
		Email:       req.Email,
		DisplayName: req.DisplayName,
	}

	err = svc.setUserDetailToRedis(ctx, details)
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateUserDetails] svc.setUserDetailToRedis() got error")
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error(ctx, metadata, err, "[UpdateUserDetails] tx.Commit() got error")
		return err
	}

	return nil
}
