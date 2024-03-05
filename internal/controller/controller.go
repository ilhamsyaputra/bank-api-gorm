package controller

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

type Controller struct {
	NasabahController
	RekeningController
	TransaksiController
}

func InitController(service *service.Service, logger *logger.Logger) *Controller {
	nasabahController := InitNasabahController(service, logger)
	rekeningController := InitRekeningController(service, logger)
	transaksiController := InitTransaksiController(service, logger)

	return &Controller{
		NasabahController:   *nasabahController,
		RekeningController:  *rekeningController,
		TransaksiController: *transaksiController,
	}
}
