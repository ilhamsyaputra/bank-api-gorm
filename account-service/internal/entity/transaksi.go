package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	Id               uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	NoRekeningAsal   string    `gorm:"type:varchar(20);not null;foreignKey"`
	NoRekeningTujuan string    `gorm:"type:varchar(20);not null"`
	TipeTransaksi    string    `gorm:"type:char(1);not null"`
	Nominal          float64   `gorm:"type:numeric(38,2);not null"`
	WaktuTransaksi   time.Time `gorm:"type:timestamp;not null;default:NOW()"`
}
