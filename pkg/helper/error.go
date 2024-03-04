package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
)

func ControllerErrorHelper(ctx *fiber.Ctx, err error, logger *logger.Logger) error {
	if err != nil {
		logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")

		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	return nil
}

func ServiceError(err error, logger *logger.Logger) error {
	if err != nil {
		logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
	}

	return err
}
