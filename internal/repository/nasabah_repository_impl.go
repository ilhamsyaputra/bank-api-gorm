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

func (r *NasabahRepositoryImpl) DaftarNasabah(nasabah entity.Nasabah) error {
	return r.db.Create(&nasabah).Error
}

func (r *NasabahRepositoryImpl) ValidateNewUser(nasabah entity.Nasabah) *gorm.DB {
	return r.db.Where("nik = ?", nasabah.Nik).Or("no_hp = ?", nasabah.NoHp).First(&nasabah)
}

func (r *NasabahRepositoryImpl) GetNoNasabah() string {
	return r.CounterRepository.GetNoNasabah()
}

func (r *NasabahRepositoryImpl) UpdateNoNasabah(tx *gorm.DB) error {
	return r.CounterRepository.UpdateNoNasabah(tx)
}

func (r *NasabahRepositoryImpl) GetNoRekening() string {
	return r.CounterRepository.GetNoRekening()
}

func (r *NasabahRepositoryImpl) UpdateNoRekening(tx *gorm.DB) error {
	return r.CounterRepository.UpdateNoRekening(tx)
}

func (r *NasabahRepositoryImpl) DaftarRekening(rekening entity.Rekening) error {
	return r.RekeningRepository.Daftar(rekening)
}
