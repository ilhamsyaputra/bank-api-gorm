package config

import (
	"fmt"

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

	return db
}
