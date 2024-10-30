package authentication

import (
	"context"

	"github.com/arifinhermawan/probi/internal/usecase/authentication"
)

type authUseCaseProvider interface {
	LogIn(ctx context.Context, req authentication.LogInReq) (int64, string, error)
}

type Handler struct {
	auth authUseCaseProvider
}

func NewHandler(auth authUseCaseProvider) *Handler {
	return &Handler{
		auth: auth,
	}
}
