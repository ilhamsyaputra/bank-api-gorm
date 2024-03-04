package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	ValidateNewUser(nasabah entity.Nasabah) *gorm.DB
	DaftarNasabah(nasabah entity.Nasabah) error

	// rekening
	DaftarRekening(rekening entity.Rekening) error

	// counter
	GetNoNasabah() string
	UpdateNoNasabah(tx *gorm.DB) error
	GetNoRekening() string
	UpdateNoRekening(tx *gorm.DB) error
}
