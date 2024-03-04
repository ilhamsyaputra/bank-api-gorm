package repository

import "github.com/ilhamsyaputra/bank-api-gorm/internal/entity"

type TransaksiRepository interface {
	CatatTransaksi(transaksi entity.Transaksi) error
}
