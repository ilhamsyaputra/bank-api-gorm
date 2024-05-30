package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type TransaksiController struct {
	transaksiService service.TransaksiService
	logger           *logger.Logger
	tracer           trace.Tracer
}

func InitTransaksiController(service service.TransaksiService, logger *logger.Logger, tracer trace.Tracer) *TransaksiController {
	return &TransaksiController{
		transaksiService: service,
		logger:           logger,
		tracer:           tracer,
	}
}

func (c *TransaksiController) GetMutasi(ctx *fiber.Ctx) error {
	newCtx, span := c.tracer.Start(ctx.Context(), "TransaksiController/GetMutasi")
	defer span.End()

	noRekening := ctx.Params("no_rekening")

	response_ := response.Response{}

	resp, err := c.transaksiService.GetMutasi(newCtx, noRekening)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: enum.Status.Error,
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   fiber.StatusOK,
		Status: enum.Status.Success,
		Remark: "get mutasi berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
