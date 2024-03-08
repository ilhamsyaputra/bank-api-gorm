package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

type TransaksiController struct {
	transaksiService service.TransaksiService
	logger           *logger.Logger
}

func InitTransaksiController(service service.TransaksiService, logger *logger.Logger) *TransaksiController {
	return &TransaksiController{
		transaksiService: service,
	}
}

func (controller *TransaksiController) GetMutasi(ctx *fiber.Ctx) error {
	noRekening := ctx.Params("no_rekening")

	response_ := response.Response{}

	resp, err := controller.transaksiService.GetMutasi(noRekening)
	if err != nil {
		helper.ControllerError(&err, controller.logger)
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
