package config

import (
	"fmt"

	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDatabase(viper *viper.Viper, log *logger.Logger) *gorm.DB {
	var (
		DB_DATABASE = viper.GetString("DB_DATABASE")
		DB_HOST     = viper.GetString("DB_HOST")
		DB_PORT     = viper.GetInt("DB_PORT")
		DB_USER     = viper.GetString("DB_USER")
		DB_PASSWORD = viper.GetString("DB_PASSWORD")
		SSLMODE     = viper.GetString("DB_SSLMODE")
		dsn         = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE, SSLMODE)
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&entity.Nasabah{},
		&entity.Enum{},
		&entity.Rekening{},
		&entity.Transaksi{},
		&entity.Counter{},
		&entity.Otp{},
	)

	if err == nil && db.Migrator().HasTable(&entity.Enum{}) {
		err = db.First(&entity.Enum{}).Error
		if err == gorm.ErrRecordNotFound {
			enums := []*entity.Enum{
				{
					Scope:       "tipe_transaksi",
					Value:       "D",
					Description: "Tarik",
				},
				{
					Scope:       "tipe_transaksi",
					Value:       "C",
					Description: "Tabung",
				},
			}

			db.Create(enums)
		}
	}

	if db.Migrator().HasTable(&entity.Counter{}) {
		err = db.First(&entity.Counter{}).Error
		if err == gorm.ErrRecordNotFound {
			counter := []*entity.Counter{
				{
					Name:  "No Nasabah",
					Value: 0,
				},
				{
					Name:  "No Rekening",
					Value: 0,
				},
			}
			db.Create(counter)
		}
	}

	return db
}
