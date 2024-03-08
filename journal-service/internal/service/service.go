package service

import (
	"journal-service/internal/repository"
	"journal-service/pkg/logger"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	logger     *logger.Logger
	db         *gorm.DB

	JournalService
}

func InitService(db *gorm.DB, repository *repository.Repository, logger *logger.Logger) *Service {
	journalService := InitJournalRepositoryImpl(db, repository.JournalRepository, validator.New(), logger)

	return &Service{
		repository: repository,
		validator:  validator.New(),
		logger:     logger,
		db:         db,

		JournalService: journalService,
	}
}
