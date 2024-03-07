package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type Repository struct {
	NasabahRepository
	RekeningRepository
	TransaksiRepository
	db  *gorm.DB
	log *logger.Logger
}

func InitRepository(db *gorm.DB, log *logger.Logger) *Repository {
	nasabahRepository := InitNasabahRepositoryImpl(db, log)
	rekeningRepository := InitRekeningRepositoryImpl(db, log)
	transaksiRepository := InitTransaksiRepositoryImpl(db, log)

	return &Repository{
		db:  db,
		log: log,

		NasabahRepository:   nasabahRepository,
		RekeningRepository:  rekeningRepository,
		TransaksiRepository: transaksiRepository,
	}
}
