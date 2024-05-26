package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	otp_ "github.com/ilhamsyaputra/bank-api-gorm/pkg/otp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type LoginRepository interface {
	CheckNoHp(context.Context, *gorm.DB, entity.Nasabah) error
	CheckOtpData(context.Context, *gorm.DB, entity.Otp) error
	GenerateOtp(context.Context, *gorm.DB, entity.Otp) error
	UpdateOtp(context.Context, *gorm.DB, entity.Otp) error
	VerifyOtp(context.Context, *gorm.DB, entity.Otp) (entity.Otp, error)
	UpdateOtpBatasCoba(context.Context, *gorm.DB, entity.Otp) error

	GetNasabah(context.Context, *gorm.DB, entity.Nasabah) (entity.Nasabah, error)
}

type LoginRepositoryImpl struct {
	logger *logger.Logger
	tracer trace.Tracer

	NasabahRepository
}

func InitLoginRepositoryImpl(db *gorm.DB, logger *logger.Logger, tracer trace.Tracer) LoginRepository {
	return &LoginRepositoryImpl{
		logger: logger,
		tracer: tracer,
	}
}

func (r *LoginRepositoryImpl) CheckNoHp(ctx context.Context, tx *gorm.DB, params entity.Nasabah) (err error) {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/Login", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", params))))
	defer span.End()

	nasabah := entity.Nasabah{}

	err = tx.Where("no_hp = ?", params.NoHp).First(&nasabah).Error
	return
}

func (r *LoginRepositoryImpl) CheckOtpData(ctx context.Context, tx *gorm.DB, otp entity.Otp) error {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/CheckOtpData", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", otp))))
	defer span.End()

	return tx.First(&otp).Error
}

func (r *LoginRepositoryImpl) GenerateOtp(ctx context.Context, tx *gorm.DB, otp entity.Otp) (err error) {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/GenerateOtp", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", otp))))
	defer span.End()

	otp.KodeOtp = strconv.Itoa(otp_.GenerateOtp())
	return tx.Create(&otp).Error
}

func (r *LoginRepositoryImpl) UpdateOtp(ctx context.Context, tx *gorm.DB, otp entity.Otp) (err error) {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/UpdateOtp", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", otp))))
	defer span.End()

	tx.First(&otp)
	otp.KodeOtp = strconv.Itoa(otp_.GenerateOtp())
	otp.BatasCoba = 5
	otp.WaktuGenerate = time.Now()
	otp.WaktuExpired = time.Now().Add(5 * time.Minute)

	return tx.Save(&otp).Error
}

func (r *LoginRepositoryImpl) VerifyOtp(ctx context.Context, tx *gorm.DB, otp entity.Otp) (otp_ entity.Otp, err error) {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/VerifyOtp", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", otp))))
	defer span.End()

	err = tx.First(&otp).Error
	otp_ = otp
	return
}

func (r *LoginRepositoryImpl) UpdateOtpBatasCoba(ctx context.Context, tx *gorm.DB, otp entity.Otp) (err error) {
	_, span := r.tracer.Start(ctx, "LoginRepositoryImpl/UpdateOtpBatasCoba", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%+v", otp))))
	defer span.End()

	tx.First(&otp)
	otp.BatasCoba -= 1
	return tx.Save(&otp).Error
}

func (r *LoginRepositoryImpl) GetNasabah(ctx context.Context, tx *gorm.DB, nasabah entity.Nasabah) (result entity.Nasabah, err error) {
	tx.Select("nama, no_nasabah, no_hp").Where("no_hp = ?", nasabah.NoHp).First(&result)
	return
}
