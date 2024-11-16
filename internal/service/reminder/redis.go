package reminder

import (
	"context"
	"fmt"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
)

const (
	redisKeyReminderList = "reminder:list:%d"
)

func (svc *Service) deleteReminderListInRedis(ctx context.Context, userID int64) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"deleteReminderListInRedis")
	defer span.End()

	key := buildRedisKeyReminderList(userID)
	err := svc.redis.Del(ctx, key)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"key":     key,
			"user_id": userID,
		}, err, "[deleteReminderListInRedis] svc.redis.Del() got error")
		return err
	}

	return nil
}

func buildRedisKeyReminderList(userID int64) string {
	return fmt.Sprintf(redisKeyReminderList, userID)
}
