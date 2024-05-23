package main

import (
	"context"
	"log"

	"github.com/ilhamsyaputra/bank-api-gorm/config"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/controller"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/server"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx := context.Background()
	viper_ := config.InitViper()

	// Service Name
	SERVICE := viper_.GetString("SERVICE")

	// tracer
	tracerProvider := config.InitTracer(ctx)
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	tracer := otel.Tracer("account-service")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	db := config.InitDatabase(viper_, logger)
	redis_ := config.InitRedis(ctx, viper_, logger)
	repository := repository.InitRepository(db, logger, tracer)
	service := service.InitService(ctx, db, repository, redis_, logger, tracer)
	controller := controller.InitController(ctx, service, logger, tracer)
	server := server.InitServer(ctx, controller)

	// Start service API
	server.Start(logger)
}
