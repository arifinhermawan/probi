package user

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/service/user"
)

func (uc *UseCase) CreateUser(ctx context.Context, req CreateUserReq) error {
	// find user by email

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
