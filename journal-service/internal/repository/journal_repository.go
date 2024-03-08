package repository

import (
	"journal-service/internal/entity"

	"gorm.io/gorm"
)

type JournalRepository interface {
	CreateJournal(tx *gorm.DB, nasabah entity.Journal) error
}
