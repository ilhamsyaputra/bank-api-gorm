package service

import (
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
)

type NasabahService interface {
	Daftar(nasabah request.DaftarRequest) (response.DaftarResponse, error)
	Login(params request.LoginRequest) (response.LoginResponse, error)
}
