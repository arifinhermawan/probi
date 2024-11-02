package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arifinhermawan/blib/log"
)

const (
	redisKeyJWT = "auth:jwt:%d"
)

func (svc *Service) deleteJWTFromRedis(ctx context.Context, userID int64) error {
	key := buildRedisJWTKey(userID)
	err := svc.redis.Del(ctx, key)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[deleteJWTInRedis] svc.redis.Del() got error")
		return err
	}

	return nil
}

func (svc *Service) getJWTFromRedis(ctx context.Context, userID int64) (string, error) {
	key := buildRedisJWTKey(userID)

	metadata := map[string]interface{}{
		"user_id": userID,
	}

	jwt, err := svc.redis.Get(ctx, key)
	if err != nil {
		log.Error(ctx, metadata, err, "[getJWTFromRedis] svc.redis.Get() got error")
		return "", err
	}

	if jwt == "" {
		return "", nil
	}

	var result string
	err = json.Unmarshal([]byte(jwt), &result)
	if err != nil {
		log.Error(ctx, metadata, err, "[getJWTFromRedis] svc.redis.Get() got error")
		return "", err
	}

	return result, nil
}

func (svc *Service) setJWTToRedis(ctx context.Context, userID int64, jwt string) error {
	key := buildRedisJWTKey(userID)
	ttl := time.Second * time.Duration(svc.lib.GetConfig().TTL.JWT)

	err := svc.redis.Set(ctx, key, jwt, ttl)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id":        userID,
			"ttl_in_seconds": ttl,
		}, err, "[setJWTToRedis] svc.redis.Set() got error")

		return err
	}

	return nil
}

func buildRedisJWTKey(userID int64) string {
	return fmt.Sprintf(redisKeyJWT, userID)
}
