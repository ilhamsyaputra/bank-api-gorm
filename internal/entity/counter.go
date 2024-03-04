package entity

type Counter struct {
	Name  string `gorm:"type:varchar(20);not null;primaryKey"`
	Value int64  `gorm:"type:int"`
}
