package rabbitmq

import "github.com/streadway/amqp"

type rmqPublishProvider interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

type RMQRepo struct {
	rmq rmqPublishProvider
}

func NewRMQRepo(rmq rmqPublishProvider) *RMQRepo {
	return &RMQRepo{
		rmq: rmq,
	}
}
