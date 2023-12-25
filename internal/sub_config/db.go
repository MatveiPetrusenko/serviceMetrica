package sub_config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"bdName"`
	SSLMode  string `yaml:"sslMode"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

// Validation ...
func (d *DBConfig) Validation() error {
	err := validation.ValidateStruct(
		d, //  field #0 must be specified as a pointer
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.Port, validation.Required),
	)

	return err
}
