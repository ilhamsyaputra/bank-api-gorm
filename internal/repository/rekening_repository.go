package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type RekeningRepository interface {
	Daftar(tx *gorm.DB, rekening entity.Rekening) error

	CheckRekening(tx *gorm.DB, rekening entity.Rekening) error
	UpdateSaldo(tx *gorm.DB, rekening entity.Rekening, nominal float64) error
	GetSaldo(tx *gorm.DB, rekening entity.Rekening) (float64, error)

	// transaksi
	CatatTransaksi(tx *gorm.DB, transaksi entity.Transaksi) error
}
