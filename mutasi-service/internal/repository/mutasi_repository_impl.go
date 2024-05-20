package repository

import (
	"mutasi-service/internal/entity"
	"mutasi-service/pkg/logger"

	"gorm.io/gorm"
)

type MutasiRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitMutasiRepositoryImpl(db *gorm.DB, logger *logger.Logger) MutasiRepository {
	return &MutasiRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *MutasiRepositoryImpl) CreateMutasi(tx *gorm.DB, mutasi entity.Mutasi) error {
	return tx.Create(&mutasi).Error
}
