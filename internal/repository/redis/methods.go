package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arifinhermawan/blib/log"
)

func (repo *RedisRepo) Del(ctx context.Context, key string) error {
	redisInt := repo.redis.Del(ctx, key)
	_, err := redisInt.Result()
	if err != nil {
		meta := map[string]interface{}{
			"key": key,
		}

		log.Error(ctx, meta, err, "[Del] redisInt.Result() got error")
		return err
	}

	return nil
}

func (repo *RedisRepo) Get(ctx context.Context, key string) (string, error) {
	meta := map[string]interface{}{
		"key": key,
	}

	redisInt := repo.redis.Exists(ctx, key)
	isExist, err := redisInt.Result()
	if err != nil {
		log.Error(ctx, meta, err, "[Get] repo.redis.Exists() got error")
		return "", err
	}

	if isExist == 0 {
		return "", nil
	}

	result, err := repo.redis.Get(ctx, key).Result()
	if err != nil {
		log.Error(ctx, meta, err, "[Get] repo.redis.Get().Result() got error")
		return "", err
	}

	return result, nil
}

func (repo *RedisRepo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	meta := map[string]interface{}{
		"key":        key,
		"expiration": expiration,
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		log.Error(ctx, meta, err, "[Set] json.Marshal() got error")
		return err
	}

	redisStatus := repo.redis.Set(ctx, key, bytes, expiration)
	_, err = redisStatus.Result()
	if err != nil {
		log.Error(ctx, meta, err, "[Set] repo.redis.Set() got error")
		return err
	}

	return nil
}
