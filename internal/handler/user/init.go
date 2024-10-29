package user

import (
	"context"

	"github.com/arifinhermawan/probi/internal/usecase/user"
)

type userUseCaseProvider interface {
	CreateUser(ctx context.Context, req user.CreateUserReq) error
}

type Handler struct {
	user userUseCaseProvider
}

func NewHandler(user userUseCaseProvider) *Handler {
	return &Handler{
		user: user,
	}
}
