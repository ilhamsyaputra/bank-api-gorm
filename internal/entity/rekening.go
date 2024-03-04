package entity

type Rekening struct {
	NoRekening string      `gorm:"type:varchar(20);not null;primaryKey"`
	NoNasabah  string      `gorm:"type:varchar(20);not null;foreignKey"`
	Saldo      float64     `gorm:"type:numeric(38,2);not null"`
	Transaksi  []Transaksi `gorm:"foreignKey:NoRekeningAsal"`
}
