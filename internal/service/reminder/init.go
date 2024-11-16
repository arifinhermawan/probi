package reminder

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/reminder"
)

type libProvider interface {
	ConvertToGMT7(input time.Time) time.Time
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	CreateReminderInDB(ctx context.Context, req reminder.CreateReminderReq) error
}

type redisProvider interface {
	Del(ctx context.Context, key string) error
}

type Service struct {
	lib   libProvider
	db    dbProvider
	redis redisProvider
}

func NewService(lib libProvider, db dbProvider, redis redisProvider) *Service {
	return &Service{
		lib:   lib,
		db:    db,
		redis: redis,
	}
}
