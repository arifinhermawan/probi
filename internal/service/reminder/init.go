package reminder

import (
	"context"
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
	CreateReminderInDB(ctx context.Context, req reminder.CreateReminderReq) error
	GetActiveReminderByUserIDFromDB(ctx context.Context, userID int64) ([]reminder.Reminder, error)
	GetDueReminderIDsFromDB(ctx context.Context) ([]int64, error)
	UpdateReminderInDB(ctx context.Context, req reminder.UpdateReminderReq) error
}

type nsqProvider interface {
	PublishMessageToNSQ(ctx context.Context, topic string, message []byte) error
}

type redisProvider interface {
	Del(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type Service struct {
	lib   libProvider
	db    dbProvider
	nsq   nsqProvider
	redis redisProvider
}

func NewService(lib libProvider, db dbProvider, nsq nsqProvider, redis redisProvider) *Service {
	return &Service{
		lib:   lib,
		db:    db,
		nsq:   nsq,
		redis: redis,
	}
}
