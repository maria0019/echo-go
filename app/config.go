package app

import (
	"github.com/spf13/viper"
)

const (
	defaultJwtSecret = "jwtSecret"
	defaultEnv       = "test"
)

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("env") // name of config file (without extension)
	v.SetConfigType("json")

	v.SetDefault("jwtSigningKey", defaultJwtSecret)
	v.SetDefault("env", defaultEnv)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
