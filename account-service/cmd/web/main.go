package main

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/config"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/controller"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/server"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

func main() {
	ctx := context.Background()
	viper_ := config.InitViper()

	// Service Name
	SERVICE := viper_.GetString("SERVICE")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	db := config.InitDatabase(viper_, logger)
	redis_ := config.InitRedis(ctx, viper_, logger)
	repository := repository.InitRepository(db, logger)
	service := service.InitService(ctx, db, repository, redis_, logger)
	controller := controller.InitController(ctx, service, logger)
	server := server.InitServer(controller)

	// Start service API
	server.Start(logger)
}
