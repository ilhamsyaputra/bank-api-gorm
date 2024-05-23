package controller

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	NasabahController
	RekeningController
	TransaksiController
}

func InitController(ctx context.Context, service *service.Service, logger *logger.Logger, tracer trace.Tracer) *Controller {
	nasabahController := InitNasabahController(ctx, service, logger, tracer)
	rekeningController := InitRekeningController(ctx, service, logger, tracer)
	transaksiController := InitTransaksiController(ctx, service, logger, tracer)

	return &Controller{
		NasabahController:   *nasabahController,
		RekeningController:  *rekeningController,
		TransaksiController: *transaksiController,
	}
}
