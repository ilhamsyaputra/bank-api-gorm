package repository

import "gorm.io/gorm"

type CounterRepository interface {
	GetNoNasabah() string
	UpdateNoNasabah(tx *gorm.DB) error
	GetNoRekening() string
	UpdateNoRekening(tx *gorm.DB) error
}
