package config

import (
	"context"
	"fmt"

	"mutasi-service/pkg/logger"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitRedis(ctx context.Context, viper_ *viper.Viper, logger *logger.Logger) *redis.Client {
	var (
		REDIS_HOST     = viper_.GetString("REDIS_HOST")
		REDIS_PORT     = viper_.GetString("REDIS_PORT")
		REDIS_PASSWORD = viper_.GetString("REDIS_PASSWORD")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Password: REDIS_PASSWORD,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Fatal(logrus.Fields{"error": err}, nil, "FAILED TO CONNECT TO REDIS SERVER")
		panic(err)
	}

	logger.Info(logrus.Fields{}, nil, "CONNECTED TO REDIS SERVER")

	return client
}
