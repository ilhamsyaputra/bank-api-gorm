package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

type RekeningController struct {
	rekeningService service.RekeningService
	logger          *logger.Logger
}

func InitRekeningController(service service.RekeningService, logger *logger.Logger) *RekeningController {
	return &RekeningController{
		rekeningService: service,
	}
}

func (controller *RekeningController) Tabung(ctx *fiber.Ctx) error {
	request_ := request.TabungRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	helper.ControllerErrorHelper(ctx, err, controller.logger)

	resp, err := controller.rekeningService.Tabung(request_)

	if err != nil {
		response_ = response.Response{
			Code:   http.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Remark: "transaksi tabung berhasil",
		Data:   resp,
	}

	return ctx.Status(http.StatusOK).JSON(response_)
}
