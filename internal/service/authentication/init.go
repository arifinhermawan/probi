package authentication

import (
	"context"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type redisProvider interface {
	Del(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type Service struct {
	lib   libProvider
	redis redisProvider
}

func NewService(lib libProvider, redis redisProvider) *Service {
	return &Service{
		lib:   lib,
		redis: redis,
	}
}
