package controller

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

type Controller struct {
	NasabahController
}

func InitController(service service.NasabahService, logger *logger.Logger) *Controller {
	nasabahController := InitNasabahController(service, logger)

	return &Controller{
		NasabahController: *nasabahController,
	}
}
