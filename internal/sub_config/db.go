package sub_config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DBConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSL_MODE"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
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
