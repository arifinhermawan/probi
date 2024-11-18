package reminder

import (
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/streadway/amqp"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
}

type rmqConsumerProvider interface {
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
}

type reminderUseCaseProvider interface {
}

type Handler struct {
	lib      libProvider
	reminder reminderUseCaseProvider
	rmq      rmqConsumerProvider
}

func NewHandler(lib libProvider, reminder reminderUseCaseProvider, rmq rmqConsumerProvider) *Handler {
	return &Handler{
		lib:      lib,
		reminder: reminder,
		rmq:      rmq,
	}
}
