package validator

import (
	"github.com/go-playground/validator/v10"
)

func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())

	return &Validator{validator: v}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
