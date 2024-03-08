package controller

import (
	"journal-service/internal/service"
	"journal-service/pkg/logger"
)

type Controller struct {
	JournalController
}

func InitController(service *service.Service, logger *logger.Logger) *Controller {
	journalController := InitJournalController(service, logger)

	return &Controller{
		JournalController: *journalController,
	}
}
