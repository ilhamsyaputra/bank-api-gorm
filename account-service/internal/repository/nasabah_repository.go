package repository

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	ValidateNewUser(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) *gorm.DB
	DaftarNasabah(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) error
	Login(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) (result entity.Nasabah, err error)

	// rekening
	DaftarRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error
	CheckRekening(ctx context.Context, tx *gorm.DB, rekening entity.Rekening) error

	// counter
	GetNoNasabah(ctx context.Context, tx *gorm.DB) string
	UpdateNoNasabah(ctx context.Context, tx *gorm.DB) error
	GetNoRekening(ctx context.Context, tx *gorm.DB) string
	UpdateNoRekening(ctx context.Context, tx *gorm.DB) error
}
