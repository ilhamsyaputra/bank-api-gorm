package service

import (
	"context"
	"mutasi-service/internal/data/request"
)

type MutasiService interface {
	CreateMutasi(ctx context.Context, mutasi request.CreateMutasi) error
}
