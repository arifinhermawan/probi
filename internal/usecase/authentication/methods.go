package authentication

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/lib/errors"
)

func (uc *UseCase) LogIn(ctx context.Context, req LogInReq) (int64, string, error) {
	metadata := map[string]interface{}{
		"email":    req.Email,
		"username": req.Username,
	}

	user, err := uc.getUser(ctx, req.Email, req.Username)
	if err != nil {
		log.Error(ctx, metadata, err, "[LogIn] uc.getUser() got error")
		return 0, "", err
	}

	if user.ID == 0 {
		return 0, "", errors.ErrUserNotFound
	}

	jwt, err := uc.auth.Authenticate(ctx, user.ID)
	if err != nil {
		log.Error(ctx, metadata, err, "[LogIn] uc.auth.Authenticate() got error")
		return 0, "", err
	}

	return user.ID, jwt, nil
}

func (uc *UseCase) getUser(ctx context.Context, email string, username string) (User, error) {
	var user User

	metadata := map[string]interface{}{
		"email":    email,
		"username": username,
	}

	if email != "" {
		res, err := uc.user.GetUserByEmail(ctx, email)
		if err != nil {
			log.Error(ctx, metadata, err, "[getUser] uc.user.GetUserByEmail() got error")
			return User{}, err
		}

		user = User(res)
	} else {
		res, err := uc.user.GetUserByUsername(ctx, username)
		if err != nil {
			log.Error(ctx, metadata, err, "[getUser] uc.user.GetUserByUsername() got error")
			return User{}, err
		}

		user = User(res)
	}

	return user, nil
}
