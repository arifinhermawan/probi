package server

import (
	"github.com/arifinhermawan/probi/internal/handler/http/authentication"
	"github.com/arifinhermawan/probi/internal/handler/http/reminder"
	"github.com/arifinhermawan/probi/internal/handler/http/user"
)

type Handlers struct {
	Auth     *authentication.Handler
	Reminder *reminder.Handler
	User     *user.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Auth:     authentication.NewHandler(uc.Auth),
		Reminder: reminder.NewHandler(uc.Reminder),
		User:     user.NewHandler(uc.User),
	}
}
