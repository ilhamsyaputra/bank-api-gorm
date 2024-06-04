package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	otp_ "github.com/ilhamsyaputra/bank-api-gorm/pkg/otp"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type LoginService interface {
	LoginV2(context.Context, request.LoginV2Request) error
	VerifyOtp(context.Context, request.VerifyOtpRequest) (entity.Nasabah, error)
	VerifyPin(context.Context, request.LoginRequest) (response.LoginResponse, error)
}

type LoginServiceImpl struct {
	repository repository.LoginRepository
	validate   *validator.Validate
	log        *logger.Logger
	db         *gorm.DB
	tracer     trace.Tracer
}

func InitLoginServiceImp(db *gorm.DB, repo repository.LoginRepository, validate *validator.Validate, logger *logger.Logger, tracer trace.Tracer) LoginService {
	return &LoginServiceImpl{
		repository: repo,
		db:         db,
		validate:   validate,
		log:        logger,
		tracer:     tracer,
	}
}

func (s *LoginServiceImpl) LoginV2(ctx context.Context, params request.LoginV2Request) (err error) {
	newCtx, span := s.tracer.Start(ctx, "LoginServiceImpl/LoginV2", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "LOGIN START")

	err = s.validate.Struct(params)
	if err != nil {
		s.log.Info(logrus.Fields{"error": err.Error()}, params, "ERROR on LoginV2()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	nasabah_ := entity.Nasabah{
		NoHp: params.NoHp,
	}
	err = s.repository.CheckNoHp(newCtx, tx, nasabah_)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("tidak dapat melakukan login, user tidak ditemukan")
		}
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	otp := entity.Otp{
		NoHp: params.NoHp,
	}
	err = s.repository.CheckOtpData(newCtx, tx, otp)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			err = fmt.Errorf("terjadi kesalahan")
			return
		}

		if err == gorm.ErrRecordNotFound {
			otp.KodeOtp = strconv.Itoa(otp_.GenerateOtp())
			if errs := s.repository.GenerateOtp(newCtx, tx, otp); errs != nil {
				return
			}
		}
	}

	err = s.repository.UpdateOtp(newCtx, tx, otp)

	defer s.log.Info(logrus.Fields{}, nil, "LOGIN END")

	return
}

func (s *LoginServiceImpl) VerifyOtp(ctx context.Context, params request.VerifyOtpRequest) (result entity.Nasabah, err error) {
	newCtx, span := s.tracer.Start(ctx, "LoginServiceImpl/VerifyOtp", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "LOGIN START")

	err = s.validate.Struct(params)
	if err != nil {
		s.log.Info(logrus.Fields{"error": err.Error()}, params, "ERROR on VerifyOtp()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	otp := entity.Otp{
		NoHp:    params.NoHp,
		KodeOtp: params.KodeOtp,
	}
	otpData, err := s.repository.VerifyOtp(newCtx, tx, otp)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("data tidak ditemukan, harap untuk request OTP kembali")
			return
		}
		err = fmt.Errorf("terjadi kesalahan")
		return
	}

	if otpData.WaktuExpired.Before(time.Now()) {
		err = fmt.Errorf("waktu verifikasi otp sudah habis, silahkan request OTP kembali")
		return
	}

	if otpData.BatasCoba == 0 {
		err = fmt.Errorf("sudah melewati batas coba verifikasi OTP")
		return
	}

	if otpData.KodeOtp != otp.KodeOtp {
		err = s.repository.UpdateOtpBatasCoba(newCtx, tx, otp)
		err = fmt.Errorf("verifikasi gagal, OTP tidak tepat")
		return
	}

	nasabah_ := entity.Nasabah{
		NoHp: params.NoHp,
	}
	result, err = s.repository.GetNasabah(newCtx, tx, nasabah_)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("tidak dapat melakukan login, user tidak ditemukan")
		}
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer s.log.Info(logrus.Fields{}, nil, "LOGIN END")

	return
}

func (s *LoginServiceImpl) VerifyPin(ctx context.Context, params request.LoginRequest) (result response.LoginResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "LoginServiceImpl/VerifyOtp", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "LOGIN START")

	err = s.validate.Struct(params)
	if err != nil {
		s.log.Info(logrus.Fields{"error": err.Error()}, params, "ERROR on VerifyOtp()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	result, err = NasabahService.Login(
		&NasabahServiceImpl{
			nasabahRepository: repository.InitNasabahRepositoryImpl(s.db, s.log, s.tracer),
			validate:          s.validate,
			log:               s.log,
			db:                s.db,
			tracer:            s.tracer,
		},
		newCtx,
		params,
	)

	defer s.log.Info(logrus.Fields{}, nil, "LOGIN END")

	return
}
