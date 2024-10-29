package user

import (
	"context"

	"github.com/arifinhermawan/probi/internal/service/user"
)

type authServiceProvider interface {
	GeneratePassword(password string) string
}

type userServiceProvider interface {
	CreateUser(ctx context.Context, req user.CreateUserReq) error
}

type UseCase struct {
	auth authServiceProvider
	user userServiceProvider
}

func NewUseCase(auth authServiceProvider, user userServiceProvider) *UseCase {
	return &UseCase{
		auth: auth,
		user: user,
	}
}
