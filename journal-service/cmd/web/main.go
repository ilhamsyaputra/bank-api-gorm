package main

import (
	"context"
	"journal-service/config"
	"journal-service/internal/controller"
	"journal-service/internal/repository"
	"journal-service/internal/server"
	"journal-service/internal/service"
	"journal-service/pkg/logger"
)

func main() {
	ctx := context.Background()
	viper_ := config.InitViper()

	// Service Name
	SERVICE := viper_.GetString("JOURNAL-SERVICE")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	db := config.InitDatabase(viper_, logger)

	redis_ := config.InitRedis(ctx, viper_, logger)
	repository := repository.InitRepository(db, logger)
	service := service.InitService(ctx, db, repository, redis_, logger)
	controller := controller.InitController(service, logger)
	server := server.InitServer(controller)

	// Start service API
	server.Start(logger)
}
