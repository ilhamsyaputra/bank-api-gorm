package entity

import (
	"time"

	"github.com/google/uuid"
)

type Journal struct {
	Id               uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	NoRekeningKredit string    `gorm:"type:varchar(20);not null"`
	NoRekeningDebit  string    `gorm:"type:varchar(20);not null"`
	NominalKredit    float64   `gorm:"type:numeric(38,2);not null"`
	NominalDebit     float64   `gorm:"type:numeric(38,2);not null"`
	TanggalTransaksi time.Time `gorm:"type:timestamp;not null"`
}
