package service

import (
	"mutasi-service/internal/data/request"
)

type MutasiService interface {
	CreateMutasi(mutasi request.CreateMutasi) error
}
