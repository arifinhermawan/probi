package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/service/authentication"
	"github.com/arifinhermawan/probi/internal/service/reminder"
	"github.com/arifinhermawan/probi/internal/service/user"
)

type Services struct {
	Auth     *authentication.Service
	Reminder *reminder.Service
	User     *user.Service
}

func NewService(lib *lib.Lib, repo *Repositories) *Services {
	return &Services{
		Auth:     authentication.NewService(lib, repo.Redis),
		Reminder: reminder.NewService(lib, repo.ReminderDB, repo.NSQ, repo.Redis),
		User:     user.NewService(lib, repo.UserDB, repo.Redis),
	}
}
