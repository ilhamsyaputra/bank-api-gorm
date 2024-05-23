package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type TransaksiController struct {
	ctx              context.Context
	transaksiService service.TransaksiService
	logger           *logger.Logger
	tracer           trace.Tracer
}

func InitTransaksiController(ctx context.Context, service service.TransaksiService, logger *logger.Logger, tracer trace.Tracer) *TransaksiController {
	return &TransaksiController{
		ctx:              ctx,
		transaksiService: service,
		logger:           logger,
		tracer:           tracer,
	}
}

func (controller *TransaksiController) GetMutasi(ctx *fiber.Ctx) error {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "TransaksiController/GetMutasi")
	defer span.End()

	noRekening := ctx.Params("no_rekening")

	response_ := response.Response{}

	resp, err := controller.transaksiService.GetMutasi(tracerCtx, noRekening)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   fiber.StatusOK,
		Status: "success",
		Remark: "get mutasi berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
