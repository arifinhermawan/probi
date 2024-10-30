package server

import (
	"github.com/arifinhermawan/probi/internal/handler/authentication"
	"github.com/arifinhermawan/probi/internal/handler/user"
)

type Handlers struct {
	Auth *authentication.Handler
	User *user.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Auth: authentication.NewHandler(uc.Auth),
		User: user.NewHandler(uc.User),
	}
}
