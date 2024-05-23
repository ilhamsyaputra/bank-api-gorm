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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type NasabahController struct {
	ctx            context.Context
	nasabahService service.NasabahService
	logger         *logger.Logger
	tracer         trace.Tracer
}

func InitNasabahController(ctx context.Context, service service.NasabahService, logger *logger.Logger, tracer trace.Tracer) *NasabahController {
	return &NasabahController{
		ctx:            ctx,
		nasabahService: service,
		logger:         logger,
		tracer:         tracer,
	}
}

func (controller *NasabahController) Daftar(ctx *fiber.Ctx) (err error) {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "NasabahController/Daftar", trace.WithAttributes(attribute.String("test", "test")))
	defer span.End()

	request_ := request.DaftarRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: "terjadi kesalahan pada sistem, harap hubungi technical support",
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := controller.nasabahService.Daftar(tracerCtx, request_)
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
		Code:   fiber.StatusCreated,
		Status: "success",
		Remark: "registrasi nasabah berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response_)
}

func (controller *NasabahController) Login(ctx *fiber.Ctx) (err error) {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "NasabahController/Login", trace.WithAttributes(attribute.String("test", "test")))
	defer span.End()

	request_ := request.LoginRequest{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		return
	}

	resp, err := controller.nasabahService.Login(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
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
