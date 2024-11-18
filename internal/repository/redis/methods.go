package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
)

func (repo *Repository) Del(ctx context.Context, key string) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Redis+"Del")
	defer span.End()

	redisInt := repo.redis.Del(ctx, key)
	_, err := redisInt.Result()
	if err != nil {
		meta := map[string]interface{}{
			"key": key,
		}

		log.Warn(ctx, meta, err, "[Del] redisInt.Result() got error")
		return err
	}

	return nil
}

func (repo *Repository) Get(ctx context.Context, key string) (string, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Redis+"Get")
	defer span.End()

	meta := map[string]interface{}{
		"key": key,
	}

	redisInt := repo.redis.Exists(ctx, key)
	isExist, err := redisInt.Result()
	if err != nil {
		log.Warn(ctx, meta, err, "[Get] repo.redis.Exists() got error")
		return "", err
	}

	if isExist == 0 {
		return "", nil
	}

	result, err := repo.redis.Get(ctx, key).Result()
	if err != nil {
		log.Warn(ctx, meta, err, "[Get] repo.redis.Get().Result() got error")
		return "", err
	}

	return result, nil
}

func (repo *Repository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Redis+"Set")
	defer span.End()

	meta := map[string]interface{}{
		"key":        key,
		"expiration": expiration,
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		log.Warn(ctx, meta, err, "[Set] json.Marshal() got error")
		return err
	}

	redisStatus := repo.redis.Set(ctx, key, bytes, expiration)
	_, err = redisStatus.Result()
	if err != nil {
		log.Warn(ctx, meta, err, "[Set] repo.redis.Set() got error")
		return err
	}

	return nil
}
