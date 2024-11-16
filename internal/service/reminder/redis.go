package reminder

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
)

const (
	redisKeyReminderList = "reminder:list:%d"
)

func (svc *Service) getReminderListFromRedis(ctx context.Context, userID int64) ([]Reminder, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"getReminderListFromRedis")
	defer span.End()

	key := buildRedisKeyReminderList(userID)

	metadata := map[string]interface{}{
		"key":     key,
		"user_id": userID,
	}

	res, err := svc.redis.Get(ctx, key)
	if err != nil {
		log.Warn(ctx, metadata, err, "[getReminderListFromRedis] svc.redis.Get() got error")
		return nil, err
	}

	if res == "" {
		return nil, nil
	}

	var result []Reminder
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		log.Warn(ctx, metadata, err, "[getReminderListFromRedis] json.Unmarshal() got error")
		return nil, err
	}

	return result, nil
}

func (svc *Service) deleteReminderListInRedis(ctx context.Context, userID int64) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"deleteReminderListInRedis")
	defer span.End()

	key := buildRedisKeyReminderList(userID)
	err := svc.redis.Del(ctx, key)
	if err != nil {
		log.Warn(ctx, map[string]interface{}{
			"key":     key,
			"user_id": userID,
		}, err, "[deleteReminderListInRedis] svc.redis.Del() got error")
		return err
	}

	return nil
}

func (svc *Service) setReminderListToRedis(ctx context.Context, userID int64, reminder []Reminder) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"setReminderListToRedis")
	defer span.End()

	key := buildRedisKeyReminderList(userID)
	ttl := time.Duration(svc.lib.GetConfig().TTL.FiveMinutes) * time.Second

	err := svc.redis.Set(ctx, key, reminder, ttl)
	if err != nil {
		log.Warn(ctx, map[string]interface{}{
			"key":     key,
			"user_id": userID,
			"ttl":     ttl,
		}, err, "[setReminderListToRedis] svc.redis.Set() got error")
		return err
	}

	return nil
}

func buildRedisKeyReminderList(userID int64) string {
	return fmt.Sprintf(redisKeyReminderList, userID)
}
