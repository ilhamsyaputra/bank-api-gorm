package entity

import "time"

type Otp struct {
	NoHp          string    `gorm:"type:varchar(15);not null;primaryKey"`
	KodeOtp       string    `gorm:"type:varchar(4);not null"`
	BatasCoba     int       `gorm:"type:smallint;not null"`
	WaktuGenerate time.Time `gorm:"type:timestamp;not null;default:NOW()"`
	WaktuExpired  time.Time `gorm:"type:timestamp;not null;default:NOW() + INTERVAL '5 minutes'"`
}
