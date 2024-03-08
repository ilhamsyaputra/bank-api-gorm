package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransaksiServiceImpl struct {
	transaksiRepository repository.TransaksiRepository
	validate            *validator.Validate
	log                 *logger.Logger
	db                  *gorm.DB
}

func InitTransaksiServiceImpl(db *gorm.DB, repo repository.TransaksiRepository, validator *validator.Validate, logger *logger.Logger) TransaksiService {
	return &TransaksiServiceImpl{
		validate: validator,
		log:      logger,
		db:       db,

		transaksiRepository: repo,
	}
}

func (s *TransaksiServiceImpl) GetMutasi(noRekening string) (resp []response.GetMutasi, err error) {
	s.log.Info(logrus.Fields{}, noRekening, "GET MUTASI START")

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekening := entity.Rekening{
		NoRekening: noRekening,
	}

	result, err := s.transaksiRepository.GetMutasi(tx, rekening)
	if err != nil {
		helper.ServiceError(&err, s.log)
		return
	}

	copier.Copy(&resp, &result)

	return
}
