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

func (controller *NasabahController) Daftar(ctx *fiber.Ctx) (err error) {
	request_ := request.DaftarRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(&err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: "terjadi kesalahan pada sistem, harap hubungi technical support",
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := controller.nasabahService.Daftar(request_)
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
		Code:   fiber.StatusCreated,
		Status: "success",
		Remark: "registrasi nasabah berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response_)
}

func (controller *NasabahController) Login(ctx *fiber.Ctx) (err error) {
	request_ := request.LoginRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(&err, controller.logger)
		return
	}

	resp, err := controller.nasabahService.Login(request_)
	if err != nil {
		helper.ControllerError(&err, controller.logger)
		response_ = response.Response{
			Code:   http.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Remark: "login berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
