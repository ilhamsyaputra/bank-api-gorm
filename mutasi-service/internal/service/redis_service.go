package service

import (
	"context"
	"encoding/json"

	"mutasi-service/internal/data/request"
	"mutasi-service/pkg/logger"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type RedisService interface {
	Listen(context.Context, *redis.Client, string)
}

type RedisServiceImpl struct {
	redisClient *redis.Client
	logger      *logger.Logger
	MutasiService
}

func InitRedisService(ctx context.Context, redisClient *redis.Client, mutasiService MutasiService, logger *logger.Logger) RedisService {
	return &RedisServiceImpl{
		redisClient:   redisClient,
		logger:        logger,
		MutasiService: mutasiService,
	}
}

func (s *RedisServiceImpl) Listen(ctx context.Context, redisClient *redis.Client, channel string) {
	var err error

	// subscribe to channel
	pubsub := redisClient.Subscribe(ctx, channel)
	// defer pubsub.Close()

	_, err = pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}

			mutasi := request.CreateMutasi{}
			err = json.Unmarshal([]byte(msg.Payload), &mutasi)
			if err != nil {
				panic(err)
			}
			s.logger.Info(logrus.Fields{"received_message": msg.Payload, "event": mutasi.Event}, msg.Payload, "MESSAGE RECEIVED: "+mutasi.Event)

			carrier := propagation.MapCarrier(mutasi.TraceContext)
			ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

			err = s.MutasiService.CreateMutasi(ctx, mutasi)
			if err != nil {
				s.logger.Error(logrus.Fields{"error": err}, mutasi, "ERROR on CreateMutasi")
			}
		}
	}()
}
