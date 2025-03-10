package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment            string        `mapstructure:"ENV"`
	DataScourceURL         string        `mapstructure:"DATA_SOURCE_URL"`
	ApplicationPort        string        `mapstructure:"APPLICATION_PORT"`
	HTTPServerReadTimeout  time.Duration `mapstructure:"HTTP_SERVER_READ_TIMEOUT"`
	HTTPServerWriteTimeout time.Duration `mapstructure:"HTTP_SERVER_WRITE_TIMEOUT"`
	HTTPServerIdleimeout   time.Duration `mapstructure:"HTTP_SERVER_IDLE_TIMEOUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Cannot read config: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Cannot unmarshal config")
		return
	}

	return
}
