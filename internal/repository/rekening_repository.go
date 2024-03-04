package repository

import "github.com/ilhamsyaputra/bank-api-gorm/internal/entity"

type RekeningRepository interface {
	Daftar(rekening entity.Rekening) error
}
