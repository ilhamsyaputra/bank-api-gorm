package repository

import "gorm.io/gorm"

type CounterRepository interface {
	GetNoNasabah(tx *gorm.DB) string
	UpdateNoNasabah(tx *gorm.DB) error
	GetNoRekening(tx *gorm.DB) string
	UpdateNoRekening(tx *gorm.DB) error
}
