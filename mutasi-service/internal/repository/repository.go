package repository

import (
	"mutasi-service/pkg/logger"

	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	log *logger.Logger
	MutasiRepository
}

func InitRepository(db *gorm.DB, log *logger.Logger) *Repository {
	mutasiRepository := InitMutasiRepositoryImpl(db, log)

	return &Repository{
		db:  db,
		log: log,

		MutasiRepository: mutasiRepository,
	}
}
