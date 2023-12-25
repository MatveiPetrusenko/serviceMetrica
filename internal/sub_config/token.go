package sub_config

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type JwtToken struct {
	Secret    string        `env:"JWT_SECRET"`
	Issuer    string        `env:"JWT_ISSUER"`
	ExpiresAt time.Duration `env:"JWT_EXPIRES"`
}

// Validation ...
func (j *JwtToken) Validation() error {
	err := validation.ValidateStruct(
		j, //  field #0 must be specified as a pointer
		validation.Field(&j.Secret, validation.Required),
		validation.Field(&j.Issuer, validation.Required),
		validation.Field(&j.ExpiresAt, validation.Required),
	)

	return err
}
