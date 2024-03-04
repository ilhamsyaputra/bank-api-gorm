package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type Repository struct {
	NasabahRepository
	db  *gorm.DB
	log *logger.Logger
}

func InitRepository(db *gorm.DB, log *logger.Logger) *Repository {
	nasabahRepository := InitNasabahRepositoryImpl(db, log)

	return &Repository{
		db:  db,
		log: log,

		NasabahRepository: nasabahRepository,
	}
}
