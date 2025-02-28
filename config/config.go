package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment               string        `mapstructure:"ENV"`
	DataScourceURL            string        `mapstructure:"DATA_SOURCE_URL"`
	ApplicationPort           string        `mapstructure:"APPLICATION_PORT"`
	HTTPServerReadTimeoutSec  time.Duration `mapstructure:"HTTP_SERVER_READ_TIMEOUT_SEC"`
	HTTPServerWriteTimeoutSec time.Duration `mapstructure:"HTTP_SERVER_WRITE_TIMEOUT_SEC"`
	HTTPServerIdleimeoutSec   time.Duration `mapstructure:"HTTP_SERVER_IDLE_TIMEOUT_SEC"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Cannot read config")
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Cannot unmarshal config")
		return
	}

	return
}
