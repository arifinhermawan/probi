package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
	"github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/arifinhermawan/probi/internal/service/reminder"
	"github.com/arifinhermawan/probi/internal/service/rmq"
)

type Services struct {
	RMQ      *rmq.Service
	Reminder *reminder.Service
}

func NewService(lib *lib.Lib, db *PSQL, redis *redis.RedisRepo, rabbit *rabbitmq.RMQRepo) *Services {
	return &Services{
		Reminder: reminder.NewService(lib, db.Reminder, redis),
		RMQ:      rmq.NewService(rabbit),
	}
}
