package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	logger     *logger.Logger
	db         *gorm.DB
	redis_     *redis.Client

	NasabahService
	RekeningService
	TransaksiService
}

func InitService(ctx context.Context, db *gorm.DB, repository *repository.Repository, redis_ *redis.Client, logger *logger.Logger, tracer trace.Tracer) *Service {
	nasabahService := InitNasabahRepositoryImpl(db, repository.NasabahRepository, validator.New(), logger, tracer)
	rekeningService := InitRekeningRepositoryImpl(ctx, db, repository.RekeningRepository, redis_, validator.New(), logger, tracer)
	transaksiService := InitTransaksiServiceImpl(db, repository.TransaksiRepository, validator.New(), logger, tracer)

	return &Service{
		repository: repository,
		validator:  validator.New(),
		logger:     logger,
		db:         db,
		redis_:     redis_,

		NasabahService:   nasabahService,
		RekeningService:  rekeningService,
		TransaksiService: transaksiService,
	}
}
