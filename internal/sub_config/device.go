package sub_config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type DeviceConfig struct {
	Host string `env:"HOST_DEVICE"`
	Port string `env:"PORT_DEVICE"`
}

// Validation ...
func (d *DeviceConfig) Validation() error {
	err := validation.ValidateStruct(
		d, //  field #0 must be specified as a pointer
		validation.Field(&d.Host, is.Host, validation.Required),
		validation.Field(&d.Port, is.Port, validation.Required),
	)

	return err
}
