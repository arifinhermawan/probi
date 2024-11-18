package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/usecase/authentication"
	"github.com/arifinhermawan/probi/internal/usecase/reminder"
	"github.com/arifinhermawan/probi/internal/usecase/user"
)

type UseCases struct {
	Auth     *authentication.UseCase
	Reminder *reminder.UseCase
	User     *user.UseCase
}

func NewUseCases(lib *lib.Lib, svc *Services) *UseCases {
	return &UseCases{
		Auth:     authentication.NewUseCase(svc.Auth, svc.User),
		Reminder: reminder.NewUseCase(lib, svc.Reminder, svc.RMQ),
		User:     user.NewUseCase(svc.Auth, svc.User),
	}
}
