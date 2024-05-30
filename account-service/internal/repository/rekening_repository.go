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

type RekeningRepository interface {
	Daftar(context.Context, *gorm.DB, entity.Rekening) error

	CheckRekening(context.Context, *gorm.DB, entity.Rekening) error
	UpdateSaldo(context.Context, *gorm.DB, entity.Rekening, float64) error
	GetSaldo(context.Context, *gorm.DB, entity.Rekening) (float64, error)

	// transaksi
	CatatTransaksi(context.Context, *gorm.DB, entity.Transaksi) error
}

type RekeningRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
	tracer trace.Tracer

	TransaksiRepository
}

func InitRekeningRepositoryImpl(db *gorm.DB, logger *logger.Logger, tracer trace.Tracer) RekeningRepository {
	transaksiRepository := InitTransaksiRepositoryImpl(db, logger, tracer)

	return &RekeningRepositoryImpl{
		db:     db,
		logger: logger,
		tracer: tracer,

		TransaksiRepository: transaksiRepository,
	}
}

func (r *RekeningRepositoryImpl) Daftar(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error {
	_, span := r.tracer.Start(ctx, "RekeningRepositoryImpl/Daftar", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	return tx.Create(&rekening).Error
}

func (r *RekeningRepositoryImpl) CheckRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error {
	_, span := r.tracer.Start(ctx, "RekeningRepositoryImpl/CheckRekening", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	return tx.First(&rekening).Error
}

func (r *RekeningRepositoryImpl) UpdateSaldo(ctx context.Context, tx *gorm.DB, rekening entity.Rekening, nominal float64) error {
	_, span := r.tracer.Start(ctx, "RekeningRepositoryImpl/UpdateSaldo", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	tx.First(&rekening)
	rekening.Saldo += nominal
	return r.db.Save(&rekening).Error
}

func (r *RekeningRepositoryImpl) GetSaldo(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) (saldo float64, err error) {
	_, span := r.tracer.Start(ctx, "RekeningRepositoryImpl/GetSaldo", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", rekening))))
	defer span.End()

	err = tx.First(&rekening).Error
	saldo = rekening.Saldo
	return
}

func (r *RekeningRepositoryImpl) CatatTransaksi(ctx context.Context, tx *gorm.DB, transaksi entity.Transaksi) error {
	tracerCtx, span := r.tracer.Start(ctx, "RekeningRepositoryImpl/CatatTransaksi", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", transaksi))))
	defer span.End()

	return r.TransaksiRepository.CatatTransaksi(tracerCtx, tx, transaksi)
}
