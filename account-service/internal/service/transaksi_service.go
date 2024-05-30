package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type TransaksiService interface {
	GetMutasi(context.Context, string) ([]response.GetMutasi, error)
}

type TransaksiServiceImpl struct {
	transaksiRepository repository.TransaksiRepository
	validate            *validator.Validate
	log                 *logger.Logger
	db                  *gorm.DB
	tracer              trace.Tracer
}

func InitTransaksiServiceImpl(db *gorm.DB, repo repository.TransaksiRepository, validator *validator.Validate, logger *logger.Logger, tracer trace.Tracer) TransaksiService {
	return &TransaksiServiceImpl{
		validate: validator,
		log:      logger,
		db:       db,
		tracer:   tracer,

		transaksiRepository: repo,
	}
}

func (s *TransaksiServiceImpl) GetMutasi(ctx context.Context, noRekening string) (resp []response.GetMutasi, err error) {
	newCtx, span := s.tracer.Start(ctx, "RekeningServiceImpl/GetSaldo", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", noRekening))))
	defer span.End()

	s.log.Info(logrus.Fields{}, noRekening, "GET MUTASI START")

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekening := entity.Rekening{
		NoRekening: noRekening,
	}

	result, err := s.transaksiRepository.GetMutasi(newCtx, tx, rekening)
	if err != nil {
		helper.ServiceError(err, s.log)
		return
	}

	copier.Copy(&resp, &result)

	return
}
