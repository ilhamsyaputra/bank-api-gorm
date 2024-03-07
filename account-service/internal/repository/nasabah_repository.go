package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	ValidateNewUser(tx *gorm.DB, nasabah entity.Nasabah) *gorm.DB
	DaftarNasabah(tx *gorm.DB, nasabah entity.Nasabah) error
	Login(tx *gorm.DB, nasabah entity.Nasabah) (result entity.Nasabah, err error)

	// rekening
	DaftarRekening(tx *gorm.DB, rekening entity.Rekening) error
	CheckRekening(tx *gorm.DB, rekening entity.Rekening) error

	// counter
	GetNoNasabah(tx *gorm.DB) string
	UpdateNoNasabah(tx *gorm.DB) error
	GetNoRekening(tx *gorm.DB) string
	UpdateNoRekening(tx *gorm.DB) error
}
