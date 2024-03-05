package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type TransaksiRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitTransaksiRepositoryImpl(db *gorm.DB, logger *logger.Logger) TransaksiRepository {
	return &TransaksiRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *TransaksiRepositoryImpl) CatatTransaksi(tx *gorm.DB, transaksi entity.Transaksi) error {
	return tx.Create(&transaksi).Error
}

func (r *TransaksiRepositoryImpl) GetMutasi(tx *gorm.DB, params entity.Rekening) (result []entity.Transaksi, err error) {
	err = tx.Where("no_rekening_asal = ?", params.NoRekening).
		Or("no_rekening_tujuan = ?", params.NoRekening).
		Select("waktu_transaksi", "tipe_transaksi", "nominal").
		Find(&result).Error
	return
}
