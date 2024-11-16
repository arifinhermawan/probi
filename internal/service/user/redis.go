package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
)

const (
	redisKeyUserDetail = "user:detail:%d"
)

func (svc *Service) getUserDetailFromRedis(ctx context.Context, userID int64) (User, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"getUserDetailFromRedis")
	defer span.End()

	key := buildRedisUserDetailKey(userID)
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	res, err := svc.redis.Get(ctx, key)
	if err != nil {
		log.Error(ctx, metadata, err, "[getUserDetailFromRedis] svc.redis.Get() got error")
		return User{}, err
	}

	if res == "" {
		return User{}, nil
	}

	var user User
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		log.Error(ctx, metadata, err, "[getUserDetailFromRedis] json.Unmarshal() got error")
		return User{}, err
	}

	return user, nil
}

func (svc *Service) setUserDetailToRedis(ctx context.Context, details User) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"setUserDetailToRedis")
	defer span.End()

	key := buildRedisUserDetailKey(details.ID)
	ttl := time.Second * time.Duration(svc.lib.GetConfig().TTL.FifteenMinutes)

	err := svc.redis.Set(ctx, key, details, ttl)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id":      details.ID,
			"display_name": details.DisplayName,
			"email":        details.Email,
			"username":     details.Username,
		}, err, "[setUserDetailToRedis] svc.redis.Set() got error")

		return err
	}

	return nil
}

func buildRedisUserDetailKey(userID int64) string {
	return fmt.Sprintf(redisKeyUserDetail, userID)
}
