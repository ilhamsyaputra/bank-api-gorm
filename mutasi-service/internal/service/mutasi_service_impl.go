package service

import (
	"context"
	"fmt"
	"mutasi-service/internal/data/request"
	"mutasi-service/internal/entity"
	"mutasi-service/pkg/logger"
	"time"

	"mutasi-service/internal/repository"
	Errors "mutasi-service/pkg/errors"
	"mutasi-service/pkg/helper"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type MutasiServiceImpl struct {
	mutasiRepository repository.MutasiRepository
	validate         *validator.Validate
	log              *logger.Logger
	db               *gorm.DB
	tracer           trace.Tracer
}

func InitMutasiServiceImpl(db *gorm.DB, repo repository.MutasiRepository, validator *validator.Validate, logger *logger.Logger, tracer trace.Tracer) MutasiService {
	return &MutasiServiceImpl{
		mutasiRepository: repo,
		validate:         validator,
		log:              logger,
		db:               db,
		tracer:           tracer,
	}
}

func (service *MutasiServiceImpl) CreateMutasi(ctx context.Context, mutasi request.CreateMutasi) (err error) {
	_, span := service.tracer.Start(ctx, "MutasiServiceImpl/CreateMutasi", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", mutasi))))
	defer span.End()

	service.log.Info(logrus.Fields{}, mutasi, "CREATEMUTASI START")

	err = service.validate.Struct(mutasi)
	if err != nil {
		service.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		_, errMsg := helper.RequestValidation(err)
		err = Errors.NewError(errMsg)
		service.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	tx := service.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, service.log)

	tanggalTransaksi, err := time.Parse("01-02-2006", mutasi.TanggalTransaksi)

	mutasi_ := entity.Mutasi{
		Id:               uuid.New(),
		NoRekening:       mutasi.NoRekening,
		Nominal:          mutasi.Nominal,
		JenisTransaksi:   mutasi.JenisTransaksi,
		TanggalTransaksi: tanggalTransaksi,
	}

	err = service.mutasiRepository.CreateMutasi(tx, mutasi_)
	if err != nil {
		service.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer service.log.Info(logrus.Fields{}, mutasi, "CREATEMUTASI END")

	return
}
