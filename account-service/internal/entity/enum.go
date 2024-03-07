package entity

type Enum struct {
	Scope       string      `gorm:"type:varchar(20);not null"`
	Value       string      `gorm:"type:varchar(5);not null;primaryKey"`
	Description string      `gorm:"type:varchar(20);not null"`
	Transaksi   []Transaksi `gorm:"foreignKey:TipeTransaksi"`
}
