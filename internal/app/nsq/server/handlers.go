package server

import "github.com/arifinhermawan/probi/internal/handler/mq/reminder"

type Handlers struct {
	Reminder *reminder.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Reminder: reminder.NewHandler(uc.Reminder),
	}
}
