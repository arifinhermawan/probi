package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/service/reminder"
)

type Services struct {
	Reminder *reminder.Service
}

func NewService(lib *lib.Lib, repo *Repositories) *Services {
	return &Services{
		Reminder: reminder.NewService(lib, repo.ReminderDB, repo.NSQ, repo.Redis),
	}
}
