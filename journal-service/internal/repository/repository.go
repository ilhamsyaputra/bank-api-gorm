package repository

import (
	"journal-service/pkg/logger"

	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	log *logger.Logger
	JournalRepository
}

func InitRepository(db *gorm.DB, log *logger.Logger) *Repository {
	journalRepository := InitJournalRepositoryImpl(db, log)

	return &Repository{
		db:  db,
		log: log,

		JournalRepository: journalRepository,
	}
}
