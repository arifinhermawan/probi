package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/usecase/reminder"
)

type UseCases struct {
	Reminder *reminder.UseCase
}

func NewUseCases(lib *lib.Lib, svc *Services) *UseCases {
	return &UseCases{
		Reminder: reminder.NewUseCase(lib, svc.Reminder, svc.RMQ),
	}
}
