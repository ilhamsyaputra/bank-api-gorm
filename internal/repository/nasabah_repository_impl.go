package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type NasabahRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
	CounterRepository
	RekeningRepository
}

func InitNasabahRepositoryImpl(db *gorm.DB, logger *logger.Logger) NasabahRepository {
	counterRepository := InitCounterRepositoryImpl(db, logger)
	rekeningRepository := InitRekeningRepositoryImpl(db, logger)

	return &NasabahRepositoryImpl{
		db:                 db,
		CounterRepository:  counterRepository,
		RekeningRepository: rekeningRepository,
	}
}

func (r *NasabahRepositoryImpl) DaftarNasabah(tx *gorm.DB, nasabah entity.Nasabah) error {
	return tx.Create(&nasabah).Error
}

func (r *NasabahRepositoryImpl) ValidateNewUser(tx *gorm.DB, nasabah entity.Nasabah) *gorm.DB {
	return tx.Where("nik = ?", nasabah.Nik).Or("no_hp = ?", nasabah.NoHp).First(&nasabah)
}

func (r *NasabahRepositoryImpl) GetNoNasabah(tx *gorm.DB) string {
	return r.CounterRepository.GetNoNasabah(tx)
}

func (r *NasabahRepositoryImpl) UpdateNoNasabah(tx *gorm.DB) error {
	return r.CounterRepository.UpdateNoNasabah(tx)
}

func (r *NasabahRepositoryImpl) GetNoRekening(tx *gorm.DB) string {
	return r.CounterRepository.GetNoRekening(tx)
}

func (r *NasabahRepositoryImpl) UpdateNoRekening(tx *gorm.DB) error {
	return r.CounterRepository.UpdateNoRekening(tx)
}

func (r *NasabahRepositoryImpl) DaftarRekening(tx *gorm.DB, rekening entity.Rekening) error {
	return r.RekeningRepository.Daftar(tx, rekening)
}

func (r *NasabahRepositoryImpl) CheckRekening(tx *gorm.DB, rekening entity.Rekening) error {
	return r.RekeningRepository.CheckRekening(tx, rekening)
}

func (r *NasabahRepositoryImpl) Login(tx *gorm.DB, nasabah entity.Nasabah) (result entity.Nasabah, err error) {
	err = tx.Where("no_nasabah = ?", nasabah.NoNasabah).
		Select("pin").
		First(&result).Error
	return
}
