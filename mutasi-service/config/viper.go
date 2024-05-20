package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	viper_ := viper.New()
	viper_.SetConfigName("config")
	viper_.SetConfigType("env")
	viper_.AddConfigPath(".")
	viper_.AutomaticEnv()

	err := viper_.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to init, error : %s", err))
	}
	return viper_
}
