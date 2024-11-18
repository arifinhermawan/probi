package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	mq "github.com/arifinhermawan/probi/internal/repository/nsq"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/reminder"
	cache "github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/jmoiron/sqlx"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	NSQ *mq.Repository

	// DB Repositories
	ReminderDB *reminder.Repository

	Redis *cache.Repository
}

func NewRepository(lib *lib.Lib, psql *sqlx.DB, redis *redis.Client, publisher *nsq.Producer) *Repositories {
	return &Repositories{
		NSQ:        mq.NewNSQRepo(publisher),
		ReminderDB: reminder.NewRepository(lib, psql),
		Redis:      cache.NewRedisRepository(redis),
	}
}
