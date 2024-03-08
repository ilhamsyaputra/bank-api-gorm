package main

import (
	"github.com/ilhamsyaputra/bank-api-gorm/config"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/controller"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/server"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

func main() {
	viper_ := config.InitViper()

	// Service Name
	SERVICE := viper_.GetString("SERVICE")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	db := config.InitDatabase(viper_, logger)
	repository := repository.InitRepository(db, logger)
	service := service.InitService(db, repository, logger)
	controller := controller.InitController(service, logger)
	server := server.InitServer(controller)

	// Start service API
	server.Start(logger)
}
