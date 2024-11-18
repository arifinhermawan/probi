package server

import (
	"github.com/arifinhermawan/probi/internal/handler/mq/reminder"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/streadway/amqp"
)

type Handlers struct {
	Reminder *reminder.Handler
}

func NewHandler(lib *lib.Lib, uc *UseCases, consumer *amqp.Channel) *Handlers {
	return &Handlers{
		Reminder: reminder.NewHandler(lib, uc.Reminder, consumer),
	}
}
