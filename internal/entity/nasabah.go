package entity

import "time"

type Nasabah struct {
	NoNasabah         string     `gorm:"type:varchar(20);not null;primaryKey"`
	Nama              string     `gorm:"type:varchar(60);not null"`
	Nik               string     `gorm:"type:varchar(20);not null"`
	NoHp              string     `gorm:"type:varchar(15);not null"`
	Pin               string     `gorm:"type:text;not null"`
	KodeCabang        string     `gorm:"type:varchar(5);not null"`
	TanggalRegistrasi time.Time  `gorm:"type:timestamp;not null;default:NOW()"`
	Rekening          []Rekening `gorm:"foreignKey:NoRekening"`
}
