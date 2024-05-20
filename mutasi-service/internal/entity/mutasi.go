package entity

import (
	"time"

	"github.com/google/uuid"
)

type Mutasi struct {
	Id               uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	NoRekening       string    `gorm:"type:varchar(20);not null"`
	JenisTransaksi   string    `gorm:"type:char(1);not null"`
	Nominal          float64   `gorm:"type:numeric(38,2);not null"`
	TanggalTransaksi time.Time `gorm:"type:timestamp;not null"`
}
