package config

import (
	constant "github.com/samithiwat/samithiwat-backend/src/common/constants"
	"github.com/spf13/viper"
)

type Database struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     string `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Name     string `mapstructure:"POSTGRES_DB"`
	SSL		 string `mapstructure:"POSTGRES_SSL_MODE"`
}

type Config struct {
	Database Database `mapstructure:",squash"`
	Port string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	if viper.GetString("GO_ENV") == "production" {
		viper.SetConfigName("prod")
	}


	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	if config.Port == "" {
		config.Port = constant.DefaultPort
	}
	
	return
}