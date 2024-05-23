package service

import (
	"context"
	"fmt"

	data_ "github.com/ilhamsyaputra/bank-api-gorm/internal/data/data"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type RedisService interface {
	Publish(context.Context, *redis.Client, string, interface{}) error
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

func (s *RedisServiceImpl) Publish(ctx context.Context, redisClient *redis.Client, channel string, data interface{}) (err error) {
	tracerCtx, span := s.tracer.Start(ctx, "RedisServiceImpl/Publish", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", data))))
	defer span.End()

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(tracerCtx, carrier)

	switch v := data.(type) {
	case *data_.RedisPublish:
		v.TraceContext = carrier
	case *data_.RedisPublishMutasi:
		v.TraceContext = carrier
	default:
		return fmt.Errorf("unsupported data type")
	}

	err = redisClient.Publish(ctx, channel, data).Err()
	return
}
