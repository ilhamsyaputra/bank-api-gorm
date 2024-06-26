package service

import (
	"context"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NasabahService interface {
	Daftar(context.Context, request.DaftarRequest) (response.DaftarResponse, error)
	Login(context.Context, request.LoginRequest) (response.LoginResponse, error)
}

type NasabahServiceImpl struct {
	nasabahRepository repository.NasabahRepository
	validate          *validator.Validate
	log               *logger.Logger
	db                *gorm.DB
	tracer            trace.Tracer
}

func InitNasabahRepositoryImpl(db *gorm.DB, repo repository.NasabahRepository, validator *validator.Validate, logger *logger.Logger, tracer trace.Tracer) NasabahService {
	return &NasabahServiceImpl{
		nasabahRepository: repo,
		validate:          validator,
		log:               logger,
		db:                db,
		tracer:            tracer,
	}
}

func (s *NasabahServiceImpl) Daftar(ctx context.Context, nasabah request.DaftarRequest) (resp response.DaftarResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "NasabahServiceImpl/Daftar", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", nasabah))))
	defer span.End()

	s.log.Info(logrus.Fields{}, nasabah, "Daftar Nasabah START")
	err = s.validate.Struct(nasabah)
	if err != nil {
		s.log.Info(logrus.Fields{"error": err.Error()}, nasabah, "ERROR on Daftar()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	hashedPin, err := helper.Hash(nasabah.Pin)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	nasabahModel := entity.Nasabah{
		Nama:       nasabah.Nama,
		Nik:        nasabah.Nik,
		NoHp:       nasabah.NoHp,
		Pin:        hashedPin,
		KodeCabang: nasabah.KodeCabang,
	}

	validateUser := s.nasabahRepository.ValidateNewUser(newCtx, tx, nasabahModel)
	if validateUser.RowsAffected != 0 {
		err = fmt.Errorf("tidak dapat melakukan registrasi. data nik atau no_hp sudah terdaftar di sistem")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	noNasabahCounter := s.nasabahRepository.GetNoNasabah(newCtx, tx)
	noNasabah := nasabah.KodeCabang + helper.Zfill(noNasabahCounter, "0", 6)
	nasabahModel.NoNasabah = noNasabah

	err = s.nasabahRepository.DaftarNasabah(newCtx, tx, nasabahModel)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	// registrasi rekening
	BANK_CODE_FILLER := "99"
	noRekeningCounter := s.nasabahRepository.GetNoRekening(newCtx, tx)
	noRekening := BANK_CODE_FILLER + nasabah.KodeCabang + helper.Zfill(noRekeningCounter, "0", 8)

	rekening := entity.Rekening{
		NoNasabah:  noNasabah,
		NoRekening: noRekening,
	}

	err = s.nasabahRepository.DaftarRekening(newCtx, tx, rekening)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	// update counter
	s.nasabahRepository.UpdateNoNasabah(newCtx, tx)
	s.nasabahRepository.UpdateNoRekening(newCtx, tx)

	defer s.log.Info(logrus.Fields{}, noRekening, "Daftar Nasabah END")

	return response.DaftarResponse{NoNasabah: noNasabah, NoRekening: noRekening}, nil
}

func (s *NasabahServiceImpl) Login(ctx context.Context, params request.LoginRequest) (resp response.LoginResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "NasabahServiceImpl/Login", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

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
	result, err := s.nasabahRepository.Login(newCtx, tx, nasabah_)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("tidak dapat melakukan login, user tidak ditemukan")
		}
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Pin), []byte(params.Pin))
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		err = fmt.Errorf("tidak dapat melakukan login, pin tidak tepat")
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
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		err = fmt.Errorf("terjadi kesalahan! harap hubungi technical support")
		return
	}

	resp = response.LoginResponse{
		Token: token_,
	}

	defer s.log.Info(logrus.Fields{}, resp, "LOGIN END")

	return
}
