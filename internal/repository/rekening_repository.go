package repository

import "github.com/ilhamsyaputra/bank-api-gorm/internal/entity"

type RekeningRepository interface {
	Daftar(rekening entity.Rekening) error

	CheckRekening(rekening entity.Rekening) error
	UpdateSaldo(rekening entity.Rekening, nominal float64) error
	GetSaldo(rekening entity.Rekening) (float64, error)

	// transaksi
	CatatTransaksi(transaksi entity.Transaksi) error
}
