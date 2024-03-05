package service

import "github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"

type TransaksiService interface {
	GetMutasi(params string) (resp []response.GetMutasi, err error)
}
