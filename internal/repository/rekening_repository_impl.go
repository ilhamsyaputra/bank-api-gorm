package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type RekeningRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitRekeningRepositoryImpl(db *gorm.DB, logger *logger.Logger) RekeningRepository {
	return &RekeningRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *RekeningRepositoryImpl) Daftar(rekening entity.Rekening) error {
	return r.db.Create(&rekening).Error
}
