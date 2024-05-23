package repository

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type RekeningRepository interface {
	Daftar(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error

	CheckRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error
	UpdateSaldo(ctx context.Context, tx *gorm.DB, rekening entity.Rekening, nominal float64) error
	GetSaldo(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) (float64, error)

	// transaksi
	CatatTransaksi(ctx context.Context, tx *gorm.DB, transaksi entity.Transaksi) error
}
