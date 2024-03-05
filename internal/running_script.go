package main

import (
	"fmt"

	"github.com/ilhamsyaputra/bank-api-gorm/config"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

func main() {
	viper_ := config.InitViper()

	// Service Name
	SERVICE := viper_.GetString("SERVICE")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	db := config.InitDatabase(viper_, logger)

	transaksi := []entity.Transaksi{}
	var nominalTarik float64
	var nominalSetor float64
	var nominalTransfer float64

	jumlahTransaksi := db.Find(&transaksi).RowsAffected
	db.Model(&entity.Transaksi{}).Select("sum(nominal) as nominal").Where("tipe_transaksi = ? and no_rekening_asal = no_rekening_tujuan", enum.TipeTransaksi.Debit).Scan(&nominalTarik)
	db.Model(&entity.Transaksi{}).Select("sum(nominal) as nominal").Where("tipe_transaksi = ? and no_rekening_asal = no_rekening_tujuan", enum.TipeTransaksi.Kredit).Scan(&nominalSetor)
	db.Model(&entity.Transaksi{}).Select("sum(nominal) as nominal").Where("tipe_transaksi = ? and no_rekening_asal <> no_rekening_tujuan", enum.TipeTransaksi.Debit).Scan(&nominalTransfer)

	fmt.Println("Jumlah total transaksi:", jumlahTransaksi)
	fmt.Println("Jumlah nominal tarik:", fmt.Sprintf("%.0f", nominalTarik))
	fmt.Println("Jumlah nominal setor:", fmt.Sprintf("%.0f", nominalSetor))
	fmt.Println("Jumlah nominal transfer:", fmt.Sprintf("%.0f", nominalTransfer))
}
