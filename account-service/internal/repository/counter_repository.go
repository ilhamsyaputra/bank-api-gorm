package repository

import (
	"context"

	"gorm.io/gorm"
)

type CounterRepository interface {
	GetNoNasabah(ctx context.Context, tx *gorm.DB) string
	UpdateNoNasabah(ctx context.Context, tx *gorm.DB) error
	GetNoRekening(ctx context.Context, tx *gorm.DB) string
	UpdateNoRekening(ctx context.Context, tx *gorm.DB) error
}
