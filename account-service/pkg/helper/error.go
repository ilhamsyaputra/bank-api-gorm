package helper

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
)

func ControllerError(err *error, logger *logger.Logger) *error {
	if err != nil {
		logger.Error(logrus.Fields{"error": err}, (*err).Error(), "ERROR on Controller")
	}
	return err
}

func ServiceError(err *error, logger *logger.Logger) *error {
	if err != nil {
		logger.Error(logrus.Fields{"error": err}, (*err).Error(), "ERROR on Service")
	}
	return err
}
