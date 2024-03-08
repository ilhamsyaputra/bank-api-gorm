package controller

import (
	"net/http"

	"journal-service/internal/data/request"
	"journal-service/internal/data/response"
	"journal-service/internal/service"
	"journal-service/pkg/helper"
	"journal-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type JournalController struct {
	journalService service.JournalService
	logger         *logger.Logger
}

func InitJournalController(service service.JournalService, logger *logger.Logger) *JournalController {
	return &JournalController{
		journalService: service,
	}
}

func (controller *JournalController) CreateJournal(ctx *fiber.Ctx) (err error) {
	request_ := request.CreateJournal{}
	response_ := response.Response{}

	err = ctx.BodyParser(&request_)
	if err != nil {
		helper.ControllerError(&err, controller.logger)
		response_ = response.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Remark: err.Error(),
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response_)
	}

	err = controller.journalService.CreateJournal(request_)
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
		Remark: "catat jurnal berhasil",
		Data:   request_,
	}

	return ctx.Status(http.StatusCreated).JSON(response_)
}
