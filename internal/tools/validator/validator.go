package validator

import "github.com/go-playground/validator/v10"

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{
		validator: validator,
	}
}

func (requestValidator *RequestValidator) Validate(request interface{}) error {
	return requestValidator.validator.Struct(request)
}
