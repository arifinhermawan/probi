package utils

import (
	"context"
	"fmt"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/streadway/amqp"
)

func InitRMQConn(ctx context.Context, cfg configuration.RMQConfig) (*amqp.Connection, error) {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Error(ctx, nil, err, "[InitRMQConn] amqp.Dial() got error")
		return nil, err
	}

	return conn, nil
}

func InitConsumer(ctx context.Context, conn *amqp.Connection) (*amqp.Channel, error) {
	consumer, err := conn.Channel()
	if err != nil {
		log.Error(ctx, nil, err, "[InitConsumer] failed to create consumer")
		return nil, err
	}

	return consumer, nil
}

func InitPublisher(ctx context.Context, conn *amqp.Connection) (*amqp.Channel, error) {
	publisher, err := conn.Channel()
	if err != nil {
		log.Error(ctx, nil, err, "[InitPublisher] failed to create publisher")
		return nil, err
	}

	return publisher, nil
}
