package service

import (
	"context"
	"fmt"

	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type RedisService interface {
	Publish(context.Context, context.Context, *redis.Client, string, interface{}) error
}

type RedisServiceImpl struct {
	redisClient *redis.Client
	tracer      trace.Tracer
}

func InitRedisService(ctx context.Context, redisClient *redis.Client, logger *logger.Logger, tracer trace.Tracer) RedisService {
	return &RedisServiceImpl{
		redisClient: redisClient,
		tracer:      tracer,
	}
}

func (s *RedisServiceImpl) Publish(ctx context.Context, tracerCtx context.Context, redisClient *redis.Client, channel string, data interface{}) (err error) {
	tracerCtx, span := s.tracer.Start(ctx, "RedisServiceImpl/Publish", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", data))))
	defer span.End()

	err = redisClient.Publish(ctx, channel, data).Err()
	return
}
