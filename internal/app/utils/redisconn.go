package utils

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

var (
	errInitRedisConn = errors.New("failed to init redis connection")
)

func InitRedisConn(ctx context.Context, cfg configuration.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error(ctx, nil, err, "[InitRedisConn] client.Ping().Result() got error")
		return nil, errInitRedisConn
	}

	log.Info(ctx, nil, nil, "successfully connect to redis")

	return client, nil
}
