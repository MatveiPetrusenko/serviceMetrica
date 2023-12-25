package config

import (
	"fmt"
	"log"
	"serviceMetrica/internal/sub_config"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const EnvConfigFile = "./.env"

var instanceConfig *Config

var onceConfig sync.Once // onceConfig

type Config struct {
	Service sub_config.ServiceConfig `env:"SERVICE"`
	Device  sub_config.DeviceConfig  `env:"DEVICE"`
	DB      sub_config.DBConfig      `env:"DB"`
	Token   sub_config.JwtToken      `env:"JWT_TOKEN"`
	Auth    sub_config.Auth          `env:"AUTHENTICATION"`
}

// New return New.Config
func New() *Config {
	onceConfig.Do(func() {
		instanceConfig = new(Config)
	})

	return instanceConfig
}

// Load ...
func (config *Config) Load() error {
	err := godotenv.Load(EnvConfigFile)
	if err != nil {
		log.Fatalf("Unable to load .env file: %s\n", err)
	}

	err = env.Parse(config)
	if err != nil {
		log.Fatalf("Unable to parse environment variables: %s\n", err)
	}

	fmt.Println(config)

	// Checking errors for each sub structure
	validators := []Validator{
		&config.Service,
		&config.Device,
		&config.DB,
		&config.Token,
	}

	for _, v := range validators {
		err = v.Validation()
		if err != nil {
			return err
		}
	}

	return nil
}
