package redis

type redisProvider interface {
}

type RedisRepo struct {
	redis redisProvider
}

func NewRedisRepository(redis redisProvider) *RedisRepo {
	return &RedisRepo{
		redis: redis,
	}
}
