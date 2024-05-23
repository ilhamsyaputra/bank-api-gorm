package service

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
)

type TransaksiService interface {
	GetMutasi(ctx context.Context, params string) (resp []response.GetMutasi, err error)
}
