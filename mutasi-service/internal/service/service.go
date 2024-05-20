package service

import (
	"context"
	"mutasi-service/pkg/logger"

	"mutasi-service/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	logger     *logger.Logger
	db         *gorm.DB
	redis_     *redis.Client

	MutasiService
}

func InitService(ctx context.Context, db *gorm.DB, repository *repository.Repository, redis_ *redis.Client, logger *logger.Logger) *Service {
	mutasiService := InitMutasiServiceImpl(db, repository.MutasiRepository, validator.New(), logger)
	redisService := InitRedisService(ctx, redis_, mutasiService, logger)
	redisService.Listen(ctx, redis_, "mutasi")

	return &Service{
		repository: repository,
		validator:  validator.New(),
		logger:     logger,
		db:         db,
		redis_:     redis_,

		MutasiService: mutasiService,
	}
}
