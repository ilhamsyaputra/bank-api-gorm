package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type RekeningRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger

	TransaksiRepository
}

func InitRekeningRepositoryImpl(db *gorm.DB, logger *logger.Logger) RekeningRepository {
	transaksiRepository := InitTransaksiRepositoryImpl(db, logger)

	return &RekeningRepositoryImpl{
		db:     db,
		logger: logger,

		TransaksiRepository: transaksiRepository,
	}
}

func (r *RekeningRepositoryImpl) Daftar(tx *gorm.DB, rekening entity.Rekening) error {
	return tx.Create(&rekening).Error
}

func (r *RekeningRepositoryImpl) CheckRekening(tx *gorm.DB, rekening entity.Rekening) error {
	return tx.First(&rekening).Error
}

func (r *RekeningRepositoryImpl) UpdateSaldo(tx *gorm.DB, rekening entity.Rekening, nominal float64) error {
	tx.First(&rekening)
	rekening.Saldo += nominal
	return r.db.Save(&rekening).Error
}

func (r *RekeningRepositoryImpl) GetSaldo(tx *gorm.DB, rekening entity.Rekening) (saldo float64, err error) {
	err = tx.First(&rekening).Error
	saldo = rekening.Saldo
	return
}

func (r *RekeningRepositoryImpl) CatatTransaksi(tx *gorm.DB, transaksi entity.Transaksi) error {
	return r.TransaksiRepository.CatatTransaksi(tx, transaksi)
}
