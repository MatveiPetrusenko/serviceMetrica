package sub_config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Auth struct {
	Salt string `env:"SALT"`
}

// Validation ...
func (a *Auth) Validation() error {
	err := validation.ValidateStruct(
		a, //  field #0 must be specified as a pointer
		validation.Field(&a.Salt, validation.Required),
	)

	return err
}
