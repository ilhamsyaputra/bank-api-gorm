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

type RekeningController struct {
	ctx             context.Context
	rekeningService service.RekeningService
	logger          *logger.Logger
	rediscontext    context.Context
	tracer          trace.Tracer
}

func InitRekeningController(ctx context.Context, service service.RekeningService, logger *logger.Logger, tracer trace.Tracer) *RekeningController {
	return &RekeningController{
		ctx:             ctx,
		rekeningService: service,
		rediscontext:    ctx,
		logger:          logger,
		tracer:          tracer,
	}
}

func (controller *RekeningController) Tabung(ctx *fiber.Ctx) error {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "RekeningController/Tabung")
	defer span.End()

	request_ := request.TabungRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: "terjadi kesalahan pada sistem, harap hubungi technical support",
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := controller.rekeningService.Tabung(tracerCtx, request_)
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
		Code:   fiber.StatusOK,
		Status: "success",
		Remark: "transaksi tabung berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (controller *RekeningController) Tarik(ctx *fiber.Ctx) error {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "RekeningController/Tarik")
	defer span.End()

	request_ := request.TarikRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: "terjadi kesalahan pada sistem, harap hubungi technical support",
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := controller.rekeningService.Tarik(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   fiber.StatusCreated,
		Status: "success",
		Remark: "transaksi tarik berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (controller *RekeningController) Transfer(ctx *fiber.Ctx) error {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "RekeningController/Transfer")
	defer span.End()

	request_ := request.TransaksiRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: "terjadi kesalahan pada sistem, harap hubungi technical support",
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := controller.rekeningService.Transfer(tracerCtx, request_)
	if err != nil {
		helper.ControllerError(err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(response_)
	}

	response_ = response.Response{
		Code:   fiber.StatusCreated,
		Status: "success",
		Remark: "transaksi transfer berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (controller *RekeningController) CekSaldo(ctx *fiber.Ctx) error {
	tracerCtx, span := controller.tracer.Start(controller.ctx, "RekeningController/Transfer")
	defer span.End()

	noRekening := ctx.Params("no_rekening")

	response_ := response.Response{}

	resp, err := controller.rekeningService.GetSaldo(tracerCtx, noRekening)
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
		Code:   fiber.StatusOK,
		Status: "success",
		Remark: "cek saldo berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
