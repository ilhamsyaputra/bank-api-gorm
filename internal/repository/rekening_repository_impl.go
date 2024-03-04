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

func (r *RekeningRepositoryImpl) Daftar(rekening entity.Rekening) error {
	return r.db.Create(&rekening).Error
}

func (r *RekeningRepositoryImpl) CheckRekening(rekening entity.Rekening) error {
	return r.db.First(&rekening).Error
}

func (r *RekeningRepositoryImpl) UpdateSaldo(rekening entity.Rekening, nominal float64) error {
	r.db.First(&rekening)
	rekening.Saldo += nominal
	return r.db.Save(&rekening).Error
}

func (r *RekeningRepositoryImpl) CatatTransaksi(transaksi entity.Transaksi) error {
	return r.TransaksiRepository.CatatTransaksi(transaksi)
}
