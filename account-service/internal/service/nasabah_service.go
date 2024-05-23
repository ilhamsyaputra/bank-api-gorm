package service

import (
	"context"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
)

type NasabahService interface {
	Daftar(context.Context, request.DaftarRequest) (response.DaftarResponse, error)
	Login(context.Context, request.LoginRequest) (response.LoginResponse, error)
}
