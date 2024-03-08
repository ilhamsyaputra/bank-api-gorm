package service

import (
	"context"
	"journal-service/internal/repository"
	"journal-service/pkg/logger"

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

	JournalService
}

func InitService(ctx context.Context, db *gorm.DB, repository *repository.Repository, redis_ *redis.Client, logger *logger.Logger) *Service {
	journalService := InitJournalServiceImpl(db, repository.JournalRepository, validator.New(), logger)
	redisService := InitRedisService(ctx, redis_, journalService, logger)
	redisService.Listen(ctx, redis_, "journal")

	return &Service{
		repository: repository,
		validator:  validator.New(),
		logger:     logger,
		db:         db,
		redis_:     redis_,

		JournalService: journalService,
	}
}
