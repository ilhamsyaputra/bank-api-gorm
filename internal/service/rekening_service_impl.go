package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RekeningServiceImpl struct {
	rekeningRepository repository.RekeningRepository
	validate           *validator.Validate
	log                *logger.Logger
	db                 *gorm.DB
}

func InitRekeningRepositoryImpl(db *gorm.DB, repo repository.RekeningRepository, validator *validator.Validate, logger *logger.Logger) RekeningService {
	return &RekeningServiceImpl{
		validate: validator,
		log:      logger,
		db:       db,

		rekeningRepository: repo,
	}
}

func (service *RekeningServiceImpl) Tabung(rekening request.TabungRequest) (resp response.TabungResponse, err error) {
	service.log.Info(logrus.Fields{}, rekening, "TRANSAKSI TABUNG START")
	err = service.validate.Struct(rekening)
	if err != nil {
		service.log.Info(logrus.Fields{"error": err.Error()}, rekening, "ERROR on Daftar()")
		return
	}

	tx := service.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, service.log)

	rekening_ := entity.Rekening{
		NoRekening: rekening.NoRekening,
	}

	err = service.rekeningRepository.CheckRekening(tx, rekening_)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening tidak valid")
		helper.ServiceError(err, service.log)
		return
	}

	err = service.rekeningRepository.UpdateSaldo(tx, rekening_, rekening.Nominal)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	transaksi_ := entity.Transaksi{
		Id:               uuid.New(),
		NoRekeningAsal:   rekening.NoRekening,
		NoRekeningTujuan: rekening.NoRekening,
		TipeTransaksi:    enum.TipeTransaksi.Kredit,
		Nominal:          rekening.Nominal,
	}

	err = service.rekeningRepository.CatatTransaksi(tx, transaksi_)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	saldo, err := service.rekeningRepository.GetSaldo(tx, rekening_)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	defer service.log.Info(logrus.Fields{}, nil, "TRANSAKSI TABUNG END")

	return response.TabungResponse{Saldo: saldo}, nil
}

func (service *RekeningServiceImpl) Tarik(rekening request.TarikRequest) (resp response.TarikResponse, err error) {
	service.log.Info(logrus.Fields{}, rekening, "TRANSAKSI TARIK START")
	err = service.validate.Struct(rekening)
	if err != nil {
		service.log.Info(logrus.Fields{"error": err.Error()}, rekening, "ERROR on Daftar()")
		return
	}

	tx := service.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, service.log)

	rekening_ := entity.Rekening{
		NoRekening: rekening.NoRekening,
	}

	err = service.rekeningRepository.CheckRekening(tx, rekening_)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening tidak valid")
		helper.ServiceError(err, service.log)
		return
	}

	saldo, err := service.rekeningRepository.GetSaldo(tx, rekening_)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	if rekening.Nominal > saldo {
		err = fmt.Errorf("saldo tidak mencukupi untuk melakukan transaksi tarik")
		helper.ServiceError(err, service.log)
		return
	}

	err = service.rekeningRepository.UpdateSaldo(tx, rekening_, -rekening.Nominal)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	transaksi_ := entity.Transaksi{
		Id:               uuid.New(),
		NoRekeningAsal:   rekening.NoRekening,
		NoRekeningTujuan: rekening.NoRekening,
		TipeTransaksi:    enum.TipeTransaksi.Debit,
		Nominal:          rekening.Nominal,
	}

	err = service.rekeningRepository.CatatTransaksi(tx, transaksi_)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	saldo, err = service.rekeningRepository.GetSaldo(tx, rekening_)
	if err != nil {
		helper.ServiceError(err, service.log)
		return
	}

	defer service.log.Info(logrus.Fields{}, nil, "TRANSAKSI TARIK END")

	return response.TarikResponse{Saldo: saldo}, nil
}
