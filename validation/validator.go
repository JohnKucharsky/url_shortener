package validation

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v Validator) Validate(data interface{}) error {
	return validate.Struct(data)
}

func NewValidator() *Validator {
	return &Validator{
		validator: validate,
	}
}
