package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/data"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/request"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/data/response"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/repository"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/helper"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type RekeningService interface {
	// CheckRekening(rekening entity.Rekening) error
	Tabung(context.Context, request.TabungRequest) (response.TabungResponse, error)
	Tarik(context.Context, request.TarikRequest) (response.TarikResponse, error)
	Transfer(context.Context, request.TransaksiRequest) (response.TransferResponse, error)
	GetSaldo(context.Context, string) (response.GetSaldo, error)
}

type RekeningServiceImpl struct {
	rekeningRepository repository.RekeningRepository
	validate           *validator.Validate
	log                *logger.Logger
	db                 *gorm.DB
	redis_             *redis.Client
	tracer             trace.Tracer

	RedisService
}

func InitRekeningRepositoryImpl(ctx context.Context, db *gorm.DB, repo repository.RekeningRepository, redis_ *redis.Client, validator *validator.Validate, logger *logger.Logger, tracer trace.Tracer) RekeningService {
	redisService := InitRedisService(ctx, redis_, logger, tracer)

	return &RekeningServiceImpl{
		validate: validator,
		log:      logger,
		db:       db,
		redis_:   redis_,
		tracer:   tracer,

		rekeningRepository: repo,
		RedisService:       redisService,
	}
}

func (s *RekeningServiceImpl) Tabung(ctx context.Context, params request.TabungRequest) (resp response.TabungResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "RekeningServiceImpl/Tabung", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "TRANSAKSI TABUNG START")
	err = s.validate.Struct(params)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err.Error()}, params, "ERROR on Daftar()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekening_ := entity.Rekening{
		NoRekening: params.NoRekening,
	}

	err = s.rekeningRepository.CheckRekening(newCtx, tx, rekening_)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening tidak valid")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	err = s.rekeningRepository.UpdateSaldo(newCtx, tx, rekening_, params.Nominal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	transaksi_ := entity.Transaksi{
		Id:               uuid.New(),
		NoRekeningAsal:   params.NoRekening,
		NoRekeningTujuan: params.NoRekening,
		TipeTransaksi:    enum.TipeTransaksi.Kredit,
		Nominal:          params.Nominal,
	}

	err = s.rekeningRepository.CatatTransaksi(newCtx, tx, transaksi_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	saldo, err := s.rekeningRepository.GetSaldo(newCtx, tx, rekening_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	dataRedis := data.RedisPublish{
		Event:            "SETOR",
		NoRekeningDebit:  params.NoRekening,
		NoRekeningKredit: params.NoRekening,
		NominalKredit:    params.Nominal,
		NominalDebit:     params.Nominal,
		TanggalTransaksi: time.Now().Format("01-02-2006"),
	}

	err = s.RedisService.Publish(newCtx, s.redis_, "journal", &dataRedis)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	dataRedisMutasi := data.RedisPublishMutasi{
		Event:            "SETOR",
		NoRekening:       params.NoRekening,
		Nominal:          params.Nominal,
		JenisTransaksi:   enum.TipeTransaksi.Kredit,
		TanggalTransaksi: time.Now().Format("01-02-2006"),
		// TraceContext:     map[string]string{"traceparent": carrier.Get("traceparent")},
	}

	err = s.RedisService.Publish(newCtx, s.redis_, "mutasi", &dataRedisMutasi)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer s.log.Info(logrus.Fields{}, nil, "TRANSAKSI TABUNG END")

	return response.TabungResponse{Saldo: saldo}, nil
}

func (s *RekeningServiceImpl) Tarik(ctx context.Context, params request.TarikRequest) (resp response.TarikResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "RekeningServiceImpl/Tarik", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "TRANSAKSI TARIK START")
	err = s.validate.Struct(params)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err.Error()}, params, "ERROR on Daftar()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekening_ := entity.Rekening{
		NoRekening: params.NoRekening,
	}

	err = s.rekeningRepository.CheckRekening(newCtx, tx, rekening_)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening tidak valid")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	saldo, err := s.rekeningRepository.GetSaldo(newCtx, tx, rekening_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	if params.Nominal > saldo {
		err = fmt.Errorf("saldo tidak mencukupi untuk melakukan transaksi tarik")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	err = s.rekeningRepository.UpdateSaldo(newCtx, tx, rekening_, -params.Nominal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	transaksi_ := entity.Transaksi{
		Id:               uuid.New(),
		NoRekeningAsal:   params.NoRekening,
		NoRekeningTujuan: params.NoRekening,
		TipeTransaksi:    enum.TipeTransaksi.Debit,
		Nominal:          params.Nominal,
	}

	err = s.rekeningRepository.CatatTransaksi(newCtx, tx, transaksi_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	saldo, err = s.rekeningRepository.GetSaldo(newCtx, tx, rekening_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	dataRedis := data.RedisPublish{
		Event:            "TARIK",
		NoRekeningDebit:  params.NoRekening,
		NoRekeningKredit: params.NoRekening,
		NominalKredit:    params.Nominal,
		NominalDebit:     params.Nominal,
		TanggalTransaksi: time.Now().Format("01-02-2006"),
	}

	err = s.RedisService.Publish(newCtx, s.redis_, "journal", &dataRedis)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	dataRedisMutasi := data.RedisPublishMutasi{
		Event:            "TARIK",
		NoRekening:       params.NoRekening,
		Nominal:          params.Nominal,
		JenisTransaksi:   enum.TipeTransaksi.Debit,
		TanggalTransaksi: time.Now().Format("01-02-2006"),
	}

	err = s.RedisService.Publish(newCtx, s.redis_, "mutasi", &dataRedisMutasi)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer s.log.Info(logrus.Fields{}, nil, "TRANSAKSI TARIK END")

	return response.TarikResponse{Saldo: saldo}, nil
}

func (s *RekeningServiceImpl) Transfer(ctx context.Context, params request.TransaksiRequest) (resp response.TransferResponse, err error) {
	newCtx, span := s.tracer.Start(ctx, "RekeningServiceImpl/Transfer", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	s.log.Info(logrus.Fields{}, params, "TRANSAKSI TRANSFER START")
	err = s.validate.Struct(params)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err.Error()}, params, "ERROR on Daftar()")
		return
	}

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekeningAsal := entity.Rekening{
		NoRekening: params.NoRekeningAsal,
	}
	err = s.rekeningRepository.CheckRekening(newCtx, tx, rekeningAsal)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening asal tidak valid")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	rekeningTujuan := entity.Rekening{
		NoRekening: params.NoRekeningTujuan,
	}
	err = s.rekeningRepository.CheckRekening(newCtx, tx, rekeningTujuan)
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("nomor rekening tujuan tidak valid")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	saldoRekeningAsal, err := s.rekeningRepository.GetSaldo(newCtx, tx, rekeningAsal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	if params.Nominal > saldoRekeningAsal {
		err = fmt.Errorf("saldo tidak mencukupi untuk melakukan transaksi transfer")
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	err = s.rekeningRepository.UpdateSaldo(newCtx, tx, rekeningAsal, -params.Nominal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	err = s.rekeningRepository.UpdateSaldo(newCtx, tx, rekeningTujuan, params.Nominal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	transaksi_ := entity.Transaksi{
		Id:               uuid.New(),
		NoRekeningAsal:   params.NoRekeningAsal,
		NoRekeningTujuan: params.NoRekeningTujuan,
		TipeTransaksi:    enum.TipeTransaksi.Debit,
		Nominal:          params.Nominal,
	}

	err = s.rekeningRepository.CatatTransaksi(newCtx, tx, transaksi_)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	saldoRekeningAsal, err = s.rekeningRepository.GetSaldo(newCtx, tx, rekeningAsal)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer s.log.Info(logrus.Fields{}, nil, "TRANSAKSI TRANSFER END")

	dataRedis := data.RedisPublish{
		Event:            "TRANSFER",
		NoRekeningDebit:  params.NoRekeningAsal,
		NoRekeningKredit: params.NoRekeningTujuan,
		NominalKredit:    params.Nominal,
		NominalDebit:     params.Nominal,
		TanggalTransaksi: time.Now().Format("01-02-2006"),
	}

	err = s.RedisService.Publish(newCtx, s.redis_, "journal", &dataRedis)
	if err != nil {
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	resp = response.TransferResponse{Saldo: saldoRekeningAsal}

	return
}

func (s *RekeningServiceImpl) GetSaldo(ctx context.Context, noRekening string) (resp response.GetSaldo, err error) {
	newCtx, span := s.tracer.Start(ctx, "RekeningServiceImpl/GetSaldo", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", noRekening))))
	defer span.End()

	s.log.Info(logrus.Fields{}, noRekening, "GET SALDO START")

	tx := s.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, s.log)

	rekeningAsal := entity.Rekening{
		NoRekening: noRekening,
	}

	saldo, err := s.rekeningRepository.GetSaldo(newCtx, tx, rekeningAsal)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("nomor rekening tidak valid")
		}
		s.log.Error(logrus.Fields{"error": err}, err.Error(), "ERROR on Service")
		return
	}

	defer s.log.Info(logrus.Fields{}, nil, "TRANSAKSI TRANSFER END")

	resp = response.GetSaldo{Saldo: saldo}

	return
}
