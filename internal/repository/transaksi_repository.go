package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CatatTransaksi(tx *gorm.DB, transaksi entity.Transaksi) error
}
