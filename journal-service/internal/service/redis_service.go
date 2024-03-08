package service

import (
	"context"
	"encoding/json"
	"fmt"

	"journal-service/internal/data/request"
	"journal-service/pkg/logger"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisService interface {
	Listen(context.Context, *redis.Client, string)
}

type RedisServiceImpl struct {
	redisClient *redis.Client
	logger      *logger.Logger
	JournalService
}

func InitRedisService(ctx context.Context, redisClient *redis.Client, journalService JournalService, logger *logger.Logger) RedisService {
	return &RedisServiceImpl{
		redisClient:    redisClient,
		logger:         logger,
		JournalService: journalService,
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

			journal := request.CreateJournal{}
			err = json.Unmarshal([]byte(msg.Payload), &journal)
			if err != nil {
				panic(err)
			}
			s.logger.Info(logrus.Fields{"received_message": msg.Payload, "event": journal.Event}, msg.Payload, "MESSAGE RECEIVED: "+journal.Event)

			fmt.Println()
			fmt.Println("=== cek dikit bre ==")
			fmt.Printf("%+v", journal)
			fmt.Println()

			err = s.JournalService.CreateJournal(journal)
			if err != nil {
				s.logger.Error(logrus.Fields{"error": err}, journal, "ERROR on CreateJournal")
			}
		}
	}()
}
