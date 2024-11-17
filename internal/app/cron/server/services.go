package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/arifinhermawan/probi/internal/service/reminder"
)

type Services struct {
	Reminder *reminder.Service
}

func NewService(lib *lib.Lib, db *PSQL, redis *redis.RedisRepo) *Services {
	return &Services{
		Reminder: reminder.NewService(lib, db.Reminder, redis),
	}
}
