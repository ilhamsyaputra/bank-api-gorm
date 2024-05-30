package repository

import (
	"context"
	"fmt"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	ValidateNewUser(context.Context, *gorm.DB, entity.Nasabah) *gorm.DB
	DaftarNasabah(context.Context, *gorm.DB, entity.Nasabah) error
	Login(context.Context, *gorm.DB, entity.Nasabah) (entity.Nasabah, error)

	// rekening
	DaftarRekening(context.Context, *gorm.DB, entity.Rekening) error
	CheckRekening(context.Context, *gorm.DB, entity.Rekening) error

	// counter
	GetNoNasabah(context.Context, *gorm.DB) string
	UpdateNoNasabah(context.Context, *gorm.DB) error
	GetNoRekening(context.Context, *gorm.DB) string
	UpdateNoRekening(context.Context, *gorm.DB) error
}

type NasabahRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
	tracer trace.Tracer
	CounterRepository
	RekeningRepository
}

func InitNasabahRepositoryImpl(db *gorm.DB, logger *logger.Logger, tracer trace.Tracer) NasabahRepository {
	counterRepository := InitCounterRepositoryImpl(db, logger, tracer)
	rekeningRepository := InitRekeningRepositoryImpl(db, logger, tracer)

	return &NasabahRepositoryImpl{
		db:                 db,
		tracer:             tracer,
		CounterRepository:  counterRepository,
		RekeningRepository: rekeningRepository,
	}
}

func (r *NasabahRepositoryImpl) DaftarNasabah(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) error {
	_, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/DaftarNasabah", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", nasabah))))
	defer span.End()

	return tx.Create(&nasabah).Error
}

func (r *NasabahRepositoryImpl) ValidateNewUser(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) *gorm.DB {
	_, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/ValidateNewUser", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", nasabah))))
	defer span.End()

	return tx.Where("nik = ?", nasabah.Nik).Or("no_hp = ?", nasabah.NoHp).First(&nasabah)
}

func (r *NasabahRepositoryImpl) GetNoNasabah(ctx context.Context, tx *gorm.DB) string {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/GetNoNasabah")
	defer span.End()

	return r.CounterRepository.GetNoNasabah(tracerCtx, tx)
}

func (r *NasabahRepositoryImpl) UpdateNoNasabah(ctx context.Context, tx *gorm.DB) error {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/UpdateNoNasabah")
	defer span.End()

	return r.CounterRepository.UpdateNoNasabah(tracerCtx, tx)
}

func (r *NasabahRepositoryImpl) GetNoRekening(ctx context.Context, tx *gorm.DB) string {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/GetNoRekening")
	defer span.End()

	return r.CounterRepository.GetNoRekening(tracerCtx, tx)
}

func (r *NasabahRepositoryImpl) UpdateNoRekening(ctx context.Context, tx *gorm.DB) error {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/UpdateNoRekening")
	defer span.End()

	return r.CounterRepository.UpdateNoRekening(tracerCtx, tx)
}

func (r *NasabahRepositoryImpl) DaftarRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/DaftarRekening", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	return r.RekeningRepository.Daftar(tracerCtx, tx, rekening)
}

func (r *NasabahRepositoryImpl) CheckRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error {
	tracerCtx, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/CheckRekening", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	return r.RekeningRepository.CheckRekening(tracerCtx, tx, rekening)
}

func (r *NasabahRepositoryImpl) Login(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) (result entity.Nasabah, err error) {
	_, span := r.tracer.Start(ctx, "NasabahRepositoryImpl/Login", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", nasabah))))
	defer span.End()

	err = tx.Where("no_nasabah = ?", nasabah.NoNasabah).
		Select("pin").
		First(&result).Error
	return
}
