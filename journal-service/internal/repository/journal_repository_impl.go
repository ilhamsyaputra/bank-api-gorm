package repository

import (
	"journal-service/internal/entity"
	"journal-service/pkg/logger"

	"gorm.io/gorm"
)

type JournalRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitJournalRepositoryImpl(db *gorm.DB, logger *logger.Logger) JournalRepository {
	return &JournalRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *JournalRepositoryImpl) CreateJournal(tx *gorm.DB, journal entity.Journal) error {
	return tx.Create(&journal).Error
}
