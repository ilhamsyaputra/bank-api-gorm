package controller

import (
	"net/http"

	"mutasi-service/internal/data/request"
	"mutasi-service/internal/data/response"
	"mutasi-service/internal/service"
	"mutasi-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MutasiController struct {
	mutasiService service.MutasiService
	logger        *logger.Logger
}

func InitMutasiController(service service.MutasiService, logger *logger.Logger) *MutasiController {
	return &MutasiController{
		mutasiService: service,
	}
}

func (c *MutasiController) CreateMutasi(ctx *fiber.Ctx) (err error) {
	request_ := request.CreateMutasi{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	err = c.mutasiService.CreateMutasi(ctx.Context(), request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

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
		Remark: "catat jurnal berhasil",
		Data:   request_,
	}

	return ctx.Status(http.StatusCreated).JSON(response_)
}
