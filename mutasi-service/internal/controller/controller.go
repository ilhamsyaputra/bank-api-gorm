package controller

import (
	"mutasi-service/internal/service"
	"mutasi-service/pkg/logger"
)

type Controller struct {
	MutasiController
}

func InitController(service *service.Service, logger *logger.Logger) *Controller {
	mutasiController := InitMutasiController(service, logger)

	return &Controller{
		MutasiController: *mutasiController,
	}
}
