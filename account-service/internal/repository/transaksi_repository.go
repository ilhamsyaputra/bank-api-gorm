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

type TransaksiRepository interface {
	CatatTransaksi(context.Context, *gorm.DB, entity.Transaksi) error
	GetMutasi(context.Context, *gorm.DB, entity.Rekening) ([]entity.Transaksi, error)
}

type TransaksiRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
	tracer trace.Tracer
}

func InitTransaksiRepositoryImpl(db *gorm.DB, logger *logger.Logger, tracer trace.Tracer) TransaksiRepository {
	return &TransaksiRepositoryImpl{
		db:     db,
		logger: logger,
		tracer: tracer,
	}
}

func (r *TransaksiRepositoryImpl) CatatTransaksi(ctx context.Context, tx *gorm.DB, transaksi entity.Transaksi) error {
	_, span := r.tracer.Start(ctx, "TransaksiRepositoryImpl/CatatTransaksi", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", transaksi))))
	defer span.End()

	return tx.Create(&transaksi).Error
}

func (r *TransaksiRepositoryImpl) GetMutasi(ctx context.Context, tx *gorm.DB, params entity.Rekening) (result []entity.Transaksi, err error) {
	_, span := r.tracer.Start(ctx, "TransaksiRepositoryImpl/GetMutasi", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	err = tx.Where("no_rekening_asal = ?", params.NoRekening).
		Or("no_rekening_tujuan = ?", params.NoRekening).
		Select("waktu_transaksi", "tipe_transaksi", "nominal").
		Find(&result).Error
	return
}
