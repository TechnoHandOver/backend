package validator

import (
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	validator_ := validator.New()
	//validator_.RegisterCustomTypeFunc(timestamps.ValidateDayOfWeek, timestamps.DayOfWeek{})
	validator_.RegisterCustomTypeFunc(timestamps.ValidateTime, timestamps.Time{})

	return &RequestValidator{
		validator: validator_,
	}
}

func (requestValidator *RequestValidator) Validate(request interface{}) error {
	return requestValidator.validator.Struct(request)
}
