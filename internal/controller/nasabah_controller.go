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

type NasabahController struct {
	nasabahService service.NasabahService
	logger         *logger.Logger
}

func InitNasabahController(service service.NasabahService, logger *logger.Logger) *NasabahController {
	return &NasabahController{
		nasabahService: service,
	}
}

func (controller *NasabahController) Daftar(ctx *fiber.Ctx) error {
	request_ := request.DaftarRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	helper.ControllerErrorHelper(ctx, err, controller.logger)

	resp, err := controller.nasabahService.Daftar(request_)

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
		Remark: "registrasi nasabah berhasil",
		Data:   resp,
	}

	return ctx.Status(http.StatusCreated).JSON(response_)
}
