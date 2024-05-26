package controller

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type LoginController struct {
	service service.LoginService
	ctx     context.Context
	logger  *logger.Logger
	tracer  trace.Tracer
}

func InitLoginController(ctx context.Context, service service.LoginService, logger *logger.Logger, tracer trace.Tracer) *LoginController {
	return &LoginController{
		service: service,
		ctx:     ctx,
		logger:  logger,
		tracer:  tracer,
	}
}

func (c *LoginController) LoginV2(ctx *fiber.Ctx) (err error) {
	tracerCtx, span := c.tracer.Start(c.ctx, "LoginController/Login")
	defer span.End()

	request_ := request.LoginV2Request{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
		return
	}

	err = c.service.LoginV2(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
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
		Remark: "request otp berhasil",
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *LoginController) VerifyOtp(ctx *fiber.Ctx) (err error) {
	tracerCtx, span := c.tracer.Start(c.ctx, "LoginController/VerifyOtp")
	defer span.End()

	request_ := request.VerifyOtpRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
		return
	}

	result, err := c.service.VerifyOtp(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
		response_ = response.Response{
			Code:   http.StatusBadRequest,
			Status: "error",
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
		Status: "success",
		Remark: "verifikasi otp berhasil",
		Data:   responseData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *LoginController) VerifyPin(ctx *fiber.Ctx) (err error) {
	tracerCtx, span := c.tracer.Start(c.ctx, "LoginController/VerifyPin")
	defer span.End()

	request_ := request.LoginRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
		return
	}

	result, err := c.service.VerifyPin(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, c.logger)
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
		Remark: "verifikasi pin berhasil",
		Data:   result,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
