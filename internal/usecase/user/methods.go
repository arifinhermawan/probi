package user

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/lib/errors"
	"github.com/arifinhermawan/probi/internal/service/user"
)

func (uc *UseCase) CreateUser(ctx context.Context, req CreateUserReq) error {
	err := uc.user.CreateUser(ctx, user.CreateUserReq{
		Email:       req.Email,
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Password:    uc.auth.GeneratePassword(req.Password),
	})
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"display_name": req.DisplayName,
			"email":        req.Email,
			"username":     req.Username,
		}, err, "[CreateUser] uc.user.CreateUser() got error")
		return err
	}

	return nil
}

func (uc *UseCase) GetUserDetails(ctx context.Context, userID int64) (User, error) {
	res, err := uc.user.GetUserByID(ctx, userID)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[GetUserDetails] uc.user.GetUserByID() got error")
	}

	if res.ID == 0 {
		return User{}, errors.ErrUserNotFound
	}

	return User{
		ID:          res.ID,
		DisplayName: res.DisplayName,
		Email:       res.Email,
		Username:    res.Username,
	}, nil
}
