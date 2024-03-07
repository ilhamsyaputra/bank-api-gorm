package service

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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

	validateUser := service.nasabahRepository.ValidateNewUser(tx, nasabahModel)
	if validateUser.RowsAffected != 0 {
		err = fmt.Errorf("tidak dapat melakukan registrasi. data nik atau no_hp sudah terdaftar di sistem")
		helper.ServiceError(err, service.log)
		return
	}

	noNasabahCounter := service.nasabahRepository.GetNoNasabah(tx)
	noNasabah := nasabah.KodeCabang + helper.Zfill(noNasabahCounter, "0", 6)
	nasabahModel.NoNasabah = noNasabah

	err = service.nasabahRepository.DaftarNasabah(tx, nasabahModel)
	helper.ServiceError(err, service.log)

	// registrasi rekening
	BANK_CODE_FILLER := "99"
	noRekeningCounter := service.nasabahRepository.GetNoRekening(tx)
	noRekening := BANK_CODE_FILLER + nasabah.KodeCabang + helper.Zfill(noRekeningCounter, "0", 8)

	rekening := entity.Rekening{
		NoNasabah:  noNasabah,
		NoRekening: noRekening,
	}

	err = service.nasabahRepository.DaftarRekening(tx, rekening)
	helper.ServiceError(err, service.log)

	// update counter
	service.nasabahRepository.UpdateNoNasabah(tx)
	service.nasabahRepository.UpdateNoRekening(tx)

	defer service.log.Info(logrus.Fields{}, noRekening, "Daftar Nasabah END")

	return response.DaftarResponse{NoRekening: noRekening}, nil
}

func (s *NasabahServiceImpl) Login(params request.LoginRequest) (resp response.LoginResponse, err error) {
	s.log.Info(logrus.Fields{}, params, "LOGIN START")
	JWT_SECRET := viper.GetString("JWT_SECRET")

	err = s.validate.Struct(params)
	if err != nil {
		s.log.Info(logrus.Fields{"error": err.Error()}, params, "ERROR on Login()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	nasabah_ := entity.Nasabah{
		NoNasabah: params.NoNasabah,
		Pin:       params.Pin,
	}
	result, err := s.nasabahRepository.Login(tx, nasabah_)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("tidak dapat melakukan login, user tidak ditemukan")
		}
		helper.ServiceError(err, s.log)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Pin), []byte(params.Pin))
	if err != nil {
		err = fmt.Errorf("tidak dapat melakukan login, pin tidak tepat")
		helper.ServiceError(err, s.log)
		return
	}

	// jwt claims
	claims := jtoken.MapClaims{
		"no_nasabah": params.NoNasabah,
		"pin":        result.Pin,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	// create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// generate encoded token
	token_, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		helper.ServiceError(err, s.log)
		err = fmt.Errorf("terjadi kesalahan! harap hubungi technical support")
		return
	}

	resp = response.LoginResponse{
		Token: token_,
	}

	defer s.log.Info(logrus.Fields{}, resp, "LOGIN END")

	return
}
