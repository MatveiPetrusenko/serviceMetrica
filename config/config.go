package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"gopkg.in/yaml.v3"
)

const EnvConfigFile = "CONFIG_PATH"

var instanceConfig *Config

var onceConfig sync.Once // onceConfig

type Config struct {
	HostInternal string `yaml:"hostInternal"`
	PortInternal string `yaml:"portInternal"`

	HostDevice string `yaml:"hostDevice"`
	PortDevice string `yaml:"portDevice"`

	JwtToken struct {
		Secret    string        `yaml:"secret"`
		Issuer    string        `yaml:"issuer"`
		ExpiresAt time.Duration `yaml:"expiresAt"`
	} `yaml:"jwtToken"`

	PostgreSQL struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"bdName"`
		SSLMode  string `yaml:"sslMode"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"postgreSQL"`
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

	return config.validation()
}

// Validation ...
func (config *Config) validation() error {
	err := validation.ValidateStruct(
		config, //  field #0 must be specified as a pointer
		validation.Field(&config.HostInternal, is.Host, validation.Required),
		validation.Field(&config.PortInternal, is.Port, validation.Required),
	)

	return err
}
