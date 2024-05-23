package service

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
)

type RekeningService interface {
	// CheckRekening(rekening entity.Rekening) error
	Tabung(ctx context.Context, params request.TabungRequest) (resp response.TabungResponse, err error)
	Tarik(ctx context.Context, params request.TarikRequest) (resp response.TarikResponse, err error)
	Transfer(ctx context.Context, params request.TransaksiRequest) (resp response.TransferResponse, err error)
	GetSaldo(ctx context.Context, params string) (resp response.GetSaldo, err error)
}
