package service

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	Publish(context.Context, *redis.Client, interface{}) error
}

type RedisServiceImpl struct {
	redisClient *redis.Client
}

func InitRedisService(ctx context.Context, redisClient *redis.Client, logger *logger.Logger) RedisService {
	return &RedisServiceImpl{
		redisClient: redisClient,
	}
}

func (s *RedisServiceImpl) Publish(ctx context.Context, redisClient *redis.Client, data interface{}) (err error) {
	err = redisClient.Publish(ctx, "journal", data).Err()
	return
}
