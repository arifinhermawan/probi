package authentication

import (
	"context"

	"github.com/arifinhermawan/probi/internal/service/user"
)

type authServiceProvider interface {
	Authenticate(ctx context.Context, userID int64) (string, error)
	GeneratePassword(password string) string
	InvalidateJWT(ctx context.Context, userID int64) error
	IsPasswordMatch(password string, encPass string) bool
}

type userServiceProvider interface {
	GetUserByEmail(ctx context.Context, email string) (user.User, error)
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
