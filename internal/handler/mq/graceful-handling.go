package mq

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/streadway/amqp"
)

func ConsumeMessages(ctx context.Context, messages <-chan amqp.Delivery, handler func(msg amqp.Delivery) error) {
	for {
		select {
		case <-ctx.Done():
			log.Info(ctx, nil, nil, "Gracefully shutdown consumer")
			return
		case msg, ok := <-messages:
			if !ok {
				log.Info(ctx, nil, nil, "Channel is closed")
				return
			}

			err := handler(msg)
			if err != nil {
				_ = msg.Nack(false, true)
				continue
			}

			msg.Ack(false)
		}
	}
}
