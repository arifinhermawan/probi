package utils

import (
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/streadway/amqp"
)

func DeclareExchange(cfg configuration.ChannelConfig, publisher *amqp.Channel) error {
	exchanges := []configuration.ExchangeConfig{
		cfg.Exchange.Reminder,
	}

	for _, exchange := range exchanges {
		err := publisher.ExchangeDeclare(exchange.Name, exchange.Type, true, false, false, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeclareQueueAndBind(cfg configuration.ChannelConfig, consumer *amqp.Channel) error {
	queues := []configuration.QueueConfig{
		cfg.Queue.ReminderSendEmail,
	}

	for _, queue := range queues {
		_, err := consumer.QueueDeclare(queue.Name, true, false, false, false, nil)
		if err != nil {
			return err
		}

		err = consumer.QueueBind(queue.Name, queue.RoutingKey, queue.Exchange, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
