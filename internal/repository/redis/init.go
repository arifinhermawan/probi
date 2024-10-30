package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisProvider interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type RedisRepo struct {
	redis redisProvider
}

func NewRedisRepository(redis redisProvider) *RedisRepo {
	return &RedisRepo{
		redis: redis,
	}
}
