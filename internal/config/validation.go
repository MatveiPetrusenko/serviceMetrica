package config

type Validator interface {
	Validation() error
}
