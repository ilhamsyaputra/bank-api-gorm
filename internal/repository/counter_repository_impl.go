package repository

import (
	"fmt"
	"strconv"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/enum"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"gorm.io/gorm"
)

type CounterRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitCounterRepositoryImpl(db *gorm.DB, logger *logger.Logger) CounterRepository {
	return &CounterRepositoryImpl{db: db, logger: logger}
}

func (r *CounterRepositoryImpl) GetNoNasabah() string {
	counter := entity.Counter{}
	r.db.Select("value").First(&counter, "name = ?", enum.Counter.NoNasabah)
	return strconv.Itoa(int(counter.Value) + 1)
}

func (r *CounterRepositoryImpl) UpdateNoNasabah(tx *gorm.DB) (err error) {
	counter := entity.Counter{}
	err = r.db.First(&counter, "name = ?", enum.Counter.NoNasabah).Error
	if err != nil {
		fmt.Println("ERROR INI PERLU DI LOG BRE")
	}

	counter.Value += 1
	err = r.db.Save(&counter).Error

	return
}

func (r *CounterRepositoryImpl) GetNoRekening() string {
	counter := entity.Counter{}
	r.db.Select("value").First(&counter, "name = ?", enum.Counter.NoRekening)
	return strconv.Itoa(int(counter.Value) + 1)
}

func (r *CounterRepositoryImpl) UpdateNoRekening(tx *gorm.DB) (err error) {
	counter := entity.Counter{}
	err = r.db.First(&counter, "name = ?", enum.Counter.NoRekening).Error
	if err != nil {
		fmt.Println("ERROR INI PERLU DI LOG BRE")
	}

	counter.Value += 1
	err = r.db.Save(&counter).Error

	return
}
