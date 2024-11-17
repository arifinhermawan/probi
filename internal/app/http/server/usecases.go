package server

import (
	"github.com/arifinhermawan/probi/internal/usecase/authentication"
	"github.com/arifinhermawan/probi/internal/usecase/reminder"
	"github.com/arifinhermawan/probi/internal/usecase/user"
)

type UseCases struct {
	Auth     *authentication.UseCase
	Reminder *reminder.UseCase
	User     *user.UseCase
}

func NewUseCases(svc *Services) *UseCases {
	return &UseCases{
		Auth:     authentication.NewUseCase(svc.Auth, svc.User),
		Reminder: reminder.NewUseCase(svc.Reminder),
		User:     user.NewUseCase(svc.Auth, svc.User),
	}
}
