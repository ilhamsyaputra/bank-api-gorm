package repository

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CatatTransaksi(ctx context.Context, tx *gorm.DB, transaksi entity.Transaksi) error
	GetMutasi(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) ([]entity.Transaksi, error)
}
