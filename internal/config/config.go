package config

import (
	"fmt"
	"os"
	"serviceMetrica/internal/sub_config"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

const EnvConfigFile = "CONFIG_PATH"

var instanceConfig *Config

var onceConfig sync.Once // onceConfig

type Config struct {
	Service *sub_config.ServiceConfig
	Device  *sub_config.DeviceConfig
	DB      *sub_config.DBConfig

	JwtToken struct {
		Secret    string        `yaml:"secret"`
		Issuer    string        `yaml:"issuer"`
		ExpiresAt time.Duration `yaml:"expiresAt"`
	} `yaml:"jwtToken"`
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
	var configPath = os.Getenv(EnvConfigFile)
	if configPath == "" {
		configPath = "config/config.yml"
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v\n", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return fmt.Errorf("Error unmarshal config file: %v\n", err)
	}

	validators := []Validator{
		config.Service,
		config.Device,
		config.DB,
	}

	for _, v := range validators {
		err = v.Validation()
		if err != nil {
			return err
		}
	}

	return nil
}
