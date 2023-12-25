package sub_config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ServiceConfig struct {
	Host string `env:"HOST_SERVICE"`
	Port string `env:"PORT_SERVICE"`
}

// Validation ...
func (s *ServiceConfig) Validation() error {
	err := validation.ValidateStruct(
		s, //  field #0 must be specified as a pointer
		validation.Field(&s.Host, is.Host, validation.Required),
		validation.Field(&s.Port, is.Port, validation.Required),
	)

	return err
}
