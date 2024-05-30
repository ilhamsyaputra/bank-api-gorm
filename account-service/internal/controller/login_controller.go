package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type LoginController struct {
	service service.LoginService
	logger  *logger.Logger
	tracer  trace.Tracer
}

func InitLoginController(service service.LoginService, logger *logger.Logger, tracer trace.Tracer) *LoginController {
	return &LoginController{
		service: service,
		logger:  logger,
		tracer:  tracer,
	}
}

func (c *LoginController) LoginV2(ctx *fiber.Ctx) (err error) {
	newCtx, span := c.tracer.Start(ctx.Context(), "LoginController/Login")
	defer span.End()

	request_ := request.LoginV2Request{}
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

	err = c.service.LoginV2(newCtx, request_)
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
		Code:   http.StatusCreated,
		Status: enum.Status.Success,
		Remark: "request otp berhasil",
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *LoginController) VerifyOtp(ctx *fiber.Ctx) (err error) {
	newCtx, span := c.tracer.Start(ctx.Context(), "LoginController/VerifyOtp")
	defer span.End()

	request_ := request.VerifyOtpRequest{}
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

	result, err := c.service.VerifyOtp(newCtx, request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")
		response_ = response.Response{
			Code:   http.StatusBadRequest,
			Status: enum.Status.Error,
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response_)
	}

	responseData := struct {
		NoNasabah string `json:"no_nasabah"`
		NoHp      string `json:"no_hp"`
		Nama      string `json:"nama"`
	}{
		NoNasabah: result.NoNasabah,
		NoHp:      result.NoHp,
		Nama:      result.Nama,
	}

	response_ = response.Response{
		Code:   http.StatusCreated,
		Status: enum.Status.Success,
		Remark: "verifikasi otp berhasil",
		Data:   responseData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *LoginController) VerifyPin(ctx *fiber.Ctx) (err error) {
	newCtx, span := c.tracer.Start(ctx.Context(), "LoginController/VerifyPin")
	defer span.End()

	request_ := request.LoginRequest{}
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

	result, err := c.service.VerifyPin(newCtx, request_)
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
		Code:   http.StatusCreated,
		Status: enum.Status.Success,
		Remark: "verifikasi pin berhasil",
		Data:   result,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
