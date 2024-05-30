package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type CounterRepository interface {
	GetNoNasabah(context.Context, *gorm.DB) string
	UpdateNoNasabah(context.Context, *gorm.DB) error
	GetNoRekening(context.Context, *gorm.DB) string
	UpdateNoRekening(context.Context, *gorm.DB) error
}

type CounterRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
	tracer trace.Tracer
}

func InitCounterRepositoryImpl(db *gorm.DB, logger *logger.Logger, tracer trace.Tracer) CounterRepository {
	return &CounterRepositoryImpl{db: db, logger: logger, tracer: tracer}
}

func (r *CounterRepositoryImpl) GetNoNasabah(ctx context.Context, tx *gorm.DB) string {
	_, span := r.tracer.Start(ctx, "CounterRepositoryImpl/GetNoNasabah")
	defer span.End()

	counter := entity.Counter{}
	tx.Select("value").First(&counter, "name = ?", enum.Counter.NoNasabah)
	return strconv.Itoa(int(counter.Value) + 1)
}

func (r *CounterRepositoryImpl) UpdateNoNasabah(ctx context.Context, tx *gorm.DB) (err error) {
	_, span := r.tracer.Start(ctx, "CounterRepositoryImpl/UpdateNoNasabah")
	defer span.End()

	counter := entity.Counter{}
	err = r.db.First(&counter, "name = ?", enum.Counter.NoNasabah).Error
	if err != nil {
		fmt.Println("ERROR INI PERLU DI LOG BRE")
	}

	counter.Value += 1
	err = r.db.Save(&counter).Error

	return
}

func (r *CounterRepositoryImpl) GetNoRekening(ctx context.Context, tx *gorm.DB) string {
	_, span := r.tracer.Start(ctx, "CounterRepositoryImpl/GetNoRekening")
	defer span.End()

	counter := entity.Counter{}
	tx.Select("value").First(&counter, "name = ?", enum.Counter.NoRekening)
	return strconv.Itoa(int(counter.Value) + 1)
}

func (r *CounterRepositoryImpl) UpdateNoRekening(ctx context.Context, tx *gorm.DB) (err error) {
	_, span := r.tracer.Start(ctx, "CounterRepositoryImpl/UpdateNoRekening")
	defer span.End()

	counter := entity.Counter{}
	err = r.db.First(&counter, "name = ?", enum.Counter.NoRekening).Error
	if err != nil {
		fmt.Println("ERROR INI PERLU DI LOG BRE")
	}

	counter.Value += 1
	err = r.db.Save(&counter).Error

	return
}
