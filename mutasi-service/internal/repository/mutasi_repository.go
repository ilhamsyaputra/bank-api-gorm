package repository

import (
	"mutasi-service/internal/entity"

	"gorm.io/gorm"
)

type MutasiRepository interface {
	CreateMutasi(tx *gorm.DB, nasabah entity.Mutasi) error
}
