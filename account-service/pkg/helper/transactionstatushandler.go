package helper

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func TransactionStatusHandler(db *gorm.DB, err *error, logger *logger.Logger) error {
	if *err != nil && (*err).Error() != "verifikasi gagal, OTP tidak tepat" {
		db.Rollback()
		logger.Error(logrus.Fields{"error": (*err).Error()}, nil, "PROCESS FAILED, ROLLBACK OCCURED")
		return *err
	}

	db.Commit()
	logger.Info(logrus.Fields{"error": nil}, nil, "PROCESS SUCCESS, COMMIT OCCURED")
	return *err
}
