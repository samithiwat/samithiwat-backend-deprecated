package config

import (
	"github.com/mitchellh/mapstructure"
	constant "github.com/samithiwat/samithiwat-backend/src/common/constants"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

//goland:noinspection ALL
type App struct {
	Port  int  `mapstructure:"port"`
	Debug bool `mapstructure:"debug"`
}

//goland:noinspection ALL
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
						env := os.Getenv(name)
						if num, err := strconv.Atoi(env); err == nil {
							result[title].(map[string]interface{})[key] = num
						} else if boolean, err := strconv.ParseBool(env); err == nil {
							result[title].(map[string]interface{})[key] = boolean
						} else {
							result[title].(map[string]interface{})[key] = env
						}
					} else {
						result[title].(map[string]interface{})[key] = temp[0]
					}
				}
				if num, ok := val.(int); ok {
					result[title].(map[string]interface{})[key] = num
				}
				if boolean, ok := val.(bool); ok {
					result[title].(map[string]interface{})[key] = boolean
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

	if config.App.Port == 0 {
		config.App.Port = constant.DefaultPort
	}

	return
}
