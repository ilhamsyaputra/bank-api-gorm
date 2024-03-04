package service

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
)

type RekeningService interface {
	// CheckRekening(rekening entity.Rekening) error
	Tabung(rekening request.TabungRequest) (resp response.TabungResponse, err error)
	Tarik(rekening request.TarikRequest) (resp response.TarikResponse, err error)
}
