package rmq

import (
	"context"

	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
)

type rabbitmqProvider interface {
	PublishMessage(ctx context.Context, req rabbitmq.PublishMessageReq) error
}

type Service struct {
	rmq rabbitmqProvider
}

func NewService(rmq rabbitmqProvider) *Service {
	return &Service{
		rmq: rmq,
	}
}
