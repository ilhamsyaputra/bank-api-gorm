package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type Services struct {
	Nasabah  NasabahService
	Rekening RekeningService
}

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	logger     *logger.Logger
	db         *gorm.DB

	NasabahService
	RekeningService
}

func InitService(db *gorm.DB, repository *repository.Repository, logger *logger.Logger) *Service {
	nasabahService := InitNasabahRepositoryImpl(db, repository.NasabahRepository, validator.New(), logger)
	rekeningService := InitRekeningRepositoryImpl(db, repository.RekeningRepository, validator.New(), logger)

	return &Service{
		repository: repository,
		validator:  validator.New(),
		logger:     logger,
		db:         db,

		NasabahService:  nasabahService,
		RekeningService: rekeningService,
	}
}
