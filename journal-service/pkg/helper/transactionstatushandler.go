package helper

import (
	"journal-service/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func TransactionStatusHandler(db *gorm.DB, err *error, logger *logger.Logger) error {
	if *err != nil {
		db.Rollback()
		logger.Error(logrus.Fields{"error": (*err).Error()}, nil, "PROCESS FAILED, ROLLBACK OCCURED")
		return *err
	}

	db.Commit()
	logger.Info(logrus.Fields{"error": nil}, nil, "PROCESS SUCCESS, COMMIT OCCURED")
	return *err
}
