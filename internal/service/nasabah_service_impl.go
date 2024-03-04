package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NasabahServiceImpl struct {
	nasabahRepository repository.NasabahRepository
	validate          *validator.Validate
	log               *logger.Logger
	db                *gorm.DB
}

func InitNasabahRepositoryImpl(db *gorm.DB, repo repository.NasabahRepository, validator *validator.Validate, logger *logger.Logger) NasabahService {
	return &NasabahServiceImpl{
		nasabahRepository: repo,
		validate:          validator,
		log:               logger,
		db:                db,
	}
}

func (service *NasabahServiceImpl) Daftar(nasabah request.DaftarRequest) (resp response.DaftarResponse, err error) {
	service.log.Info(logrus.Fields{}, nasabah, "Daftar Nasabah START")
	err = service.validate.Struct(nasabah)
	if err != nil {
		service.log.Info(logrus.Fields{"error": err.Error()}, nasabah, "ERROR on Daftar()")
		return
	}

	tx := service.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, service.log)

	hashedPin, err := helper.Hash(nasabah.Pin)
	helper.ServiceError(err, service.log)

	nasabahModel := entity.Nasabah{
		Nama:       nasabah.Nama,
		Nik:        nasabah.Nik,
		NoHp:       nasabah.NoHp,
		Pin:        hashedPin,
		KodeCabang: nasabah.KodeCabang,
	}

	validateUser := service.nasabahRepository.ValidateNewUser(nasabahModel)
	if validateUser.RowsAffected != 0 {
		err = fmt.Errorf("tidak dapat melakukan registrasi. data nik atau no_hp sudah terdaftar di sistem")
		helper.ServiceError(err, service.log)
		return
	}

	noNasabahCounter := service.nasabahRepository.GetNoNasabah()
	noNasabah := nasabah.KodeCabang + helper.Zfill(noNasabahCounter, "0", 6)
	nasabahModel.NoNasabah = noNasabah

	err = service.nasabahRepository.DaftarNasabah(nasabahModel)
	helper.ServiceError(err, service.log)

	// registrasi rekening
	BANK_CODE_FILLER := "99"
	noRekeningCounter := service.nasabahRepository.GetNoRekening()
	noRekening := BANK_CODE_FILLER + nasabah.KodeCabang + helper.Zfill(noRekeningCounter, "0", 8)

	rekening := entity.Rekening{
		NoNasabah:  noNasabah,
		NoRekening: noRekening,
	}

	err = service.nasabahRepository.DaftarRekening(rekening)
	helper.ServiceError(err, service.log)

	// update counter
	service.nasabahRepository.UpdateNoNasabah(tx)
	service.nasabahRepository.UpdateNoRekening(tx)

	defer service.log.Info(logrus.Fields{}, noRekening, "Daftar Nasabah END")

	return response.DaftarResponse{NoRekening: noRekening}, nil
}
