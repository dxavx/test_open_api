package settings

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var globalConfig *Config

func GetConfig() *Config {
	if globalConfig == nil {
		globalConfig = &Config{}
		if err := env.Parse(globalConfig); err != nil {
			log.Fatalf("config parsing error: %+v", err)
		}
	}
	return globalConfig
}

type SwaggerUI struct {
	Login    string `env:"SWAGGER_UI_LOGIN" envDefault:"swagger"`
	Password string `env:"SWAGGER_UI_PASSWORD" envDefault:"qwerty123"`
	Address  string `env:"LISTEN_ADDRESS" envDefault:"127.0.0.1:3000"`
}

type Config struct {
	SwaggerUI
}
