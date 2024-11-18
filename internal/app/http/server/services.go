package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
	"github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/arifinhermawan/probi/internal/service/authentication"
	"github.com/arifinhermawan/probi/internal/service/reminder"
	"github.com/arifinhermawan/probi/internal/service/rmq"
	"github.com/arifinhermawan/probi/internal/service/user"
)

type Services struct {
	Auth     *authentication.Service
	Reminder *reminder.Service
	RMQ      *rmq.Service
	User     *user.Service
}

func NewService(lib *lib.Lib, db *PSQL, redis *redis.RedisRepo, rabbit *rabbitmq.RMQRepo) *Services {
	return &Services{
		Auth:     authentication.NewService(lib, redis),
		Reminder: reminder.NewService(lib, db.Reminder, redis),
		User:     user.NewService(lib, db.User, redis),
		RMQ:      rmq.NewService(rabbit),
	}
}
