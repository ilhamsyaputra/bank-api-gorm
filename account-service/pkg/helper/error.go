package helper

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
)

func ControllerError(err error, logger *logger.Logger) {
	logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Controller")
}

func ServiceError(err error, logger *logger.Logger) {
	logger.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
}
