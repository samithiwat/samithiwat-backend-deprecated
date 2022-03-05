package config

import (
	"github.com/mitchellh/mapstructure"
	constant "github.com/samithiwat/samithiwat-backend/src/common/constants"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

type App struct {
	Port  string `mapstructure:"port"`
	Debug string `mapstructure:"debug"`
}

type Config struct {
	Database Database `mapstructure:",database"`
	App      App      `mapstructure:",app"`
}

func assignEnv(config *map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for title, record := range *config {
		result[title] = make(map[string]interface{})
		if rec, ok := record.(map[string]interface{}); ok {
			for key, val := range rec {
				if str, ok := val.(string); ok {
					temp := strings.Split(str, "$")
					if len(temp) > 1 {
						name := strings.Replace(temp[1], "{", "", -1)
						name = strings.Replace(name, "}", "", -1)
						result[title].(map[string]interface{})[key] = os.Getenv(name)
					}else{
						result[title].(map[string]interface{})[key] = temp[0]
					}
				}
			}
		}
	}
	return result
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if os.Getenv("GO_ENV") == "production" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("dev")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	raw := viper.AllSettings()

	raw = assignEnv(&raw)

	err = mapstructure.Decode(raw, &config)
	if err != nil {
		return
	}

	if config.App.Port == "" {
		config.App.Port = constant.DefaultPort
	}

	return
}
