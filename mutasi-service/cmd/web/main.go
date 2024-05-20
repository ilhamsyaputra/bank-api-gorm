package main

import (
	"context"
	"mutasi-service/config"
	"mutasi-service/internal/controller"
	"mutasi-service/internal/repository"
	"mutasi-service/internal/server"
	"mutasi-service/internal/service"
	"mutasi-service/pkg/logger"
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
	controller := controller.InitController(service, logger)
	server := server.InitServer(controller)

	// Start service API
	server.Start(logger)
}
