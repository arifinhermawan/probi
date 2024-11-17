package server

import (
	"github.com/arifinhermawan/probi/internal/handler/cron/reminder"
)

type Handlers struct {
	Reminder *reminder.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Reminder: reminder.NewCronHandler(uc.Reminder),
	}
}
