package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type NasabahController struct {
	nasabahService service.NasabahService
	logger         *logger.Logger
	tracer         trace.Tracer
}

func InitNasabahController(service service.NasabahService, logger *logger.Logger, tracer trace.Tracer) *NasabahController {
	return &NasabahController{
		nasabahService: service,
		logger:         logger,
		tracer:         tracer,
	}
}

func (c *NasabahController) Daftar(ctx *fiber.Ctx) (err error) {
	newCtx, span := c.tracer.Start(ctx.Context(), "NasabahController/Daftar")
	defer span.End()

	request_ := request.DaftarRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusInternalServerError,
			Status: enum.Status.Error,
			Remark: enum.Remark.InternalServerError,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := c.nasabahService.Daftar(newCtx, request_)
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
		Code:   fiber.StatusCreated,
		Status: enum.Status.Success,
		Remark: "registrasi nasabah berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response_)
}

func (c *NasabahController) Login(ctx *fiber.Ctx) (err error) {
	newCtx, span := c.tracer.Start(ctx.Context(), "NasabahController/Login")
	defer span.End()

	request_ := request.LoginRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")
		return
	}

	resp, err := c.nasabahService.Login(newCtx, request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: enum.Status.Error,
			Remark: enum.Remark.InternalServerError,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   fiber.StatusOK,
		Status: enum.Status.Success,
		Remark: "login berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
