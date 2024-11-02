package user

import (
	"context"

	"github.com/arifinhermawan/probi/internal/service/user"
)

type authServiceProvider interface {
	Authenticate(ctx context.Context, userID int64) (string, error)
	GeneratePassword(password string) string
	IsPasswordMatch(password string, encPass string) bool
}

type userServiceProvider interface {
	CreateUser(ctx context.Context, req user.CreateUserReq) error
	GetUserByEmail(ctx context.Context, email string) (user.User, error)
	GetUserByID(ctx context.Context, userID int64) (user.User, error)
	GetUserByUsername(ctx context.Context, username string) (user.User, error)
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
