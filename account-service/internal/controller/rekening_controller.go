package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/service"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type RekeningController struct {
	rekeningService service.RekeningService
	logger          *logger.Logger
	rediscontext    context.Context
	tracer          trace.Tracer
}

func InitRekeningController(ctx context.Context, service service.RekeningService, logger *logger.Logger, tracer trace.Tracer) *RekeningController {
	return &RekeningController{
		rekeningService: service,
		rediscontext:    ctx,
		logger:          logger,
		tracer:          tracer,
	}
}

func (c *RekeningController) Tabung(ctx *fiber.Ctx) error {
	newCtx, span := c.tracer.Start(ctx.Context(), "RekeningController/Tabung")
	defer span.End()

	request_ := request.TabungRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusInternalServerError,
			Status: enum.Status.Error,
			Remark: enum.Remark.InternalServerError,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := c.rekeningService.Tabung(newCtx, request_)
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
		Code:   fiber.StatusOK,
		Status: enum.Status.Success,
		Remark: "transaksi tabung berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *RekeningController) Tarik(ctx *fiber.Ctx) error {
	newCtx, span := c.tracer.Start(ctx.Context(), "RekeningController/Tarik")
	defer span.End()

	request_ := request.TarikRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusInternalServerError,
			Status: enum.Status.Error,
			Remark: enum.Remark.InternalServerError,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := c.rekeningService.Tarik(newCtx, request_)
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
		Remark: "transaksi tarik berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *RekeningController) Transfer(ctx *fiber.Ctx) error {
	newCtx, span := c.tracer.Start(ctx.Context(), "RekeningController/Transfer")
	defer span.End()

	request_ := request.TransaksiRequest{}
	response_ := response.Response{}

	err := ctx.BodyParser(&request_)
	if err != nil {
		c.logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		response_ = response.Response{
			Code:   fiber.StatusInternalServerError,
			Status: enum.Status.Error,
			Remark: enum.Remark.InternalServerError,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	resp, err := c.rekeningService.Transfer(newCtx, request_)
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
		Remark: "transaksi transfer berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}

func (c *RekeningController) CekSaldo(ctx *fiber.Ctx) error {
	newCtx, span := c.tracer.Start(ctx.Context(), "RekeningController/Transfer")
	defer span.End()

	noRekening := ctx.Params("no_rekening")

	response_ := response.Response{}

	resp, err := c.rekeningService.GetSaldo(newCtx, noRekening)
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
		Code:   fiber.StatusOK,
		Status: enum.Status.Success,
		Remark: "cek saldo berhasil",
		Data:   resp,
	}

	return ctx.Status(fiber.StatusOK).JSON(response_)
}
