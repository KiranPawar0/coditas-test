package structvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func RegisterCustomValidations(v *validator.Validate) {
	if validate == nil {
		validate = v
	}

	_ = validate.RegisterValidation("pan", func(fl validator.FieldLevel) bool {
		pattern := `^[A-Z]{5}[0-9]{4}[A-Z]$`
		match, _ := regexp.MatchString(pattern, fl.Field().String())
		return match
	})

	_ = validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		pattern := `^[0-9]{10}$`
		match, _ := regexp.MatchString(pattern, fl.Field().String())
		return match
	})
}

type StructValidator struct {
	Struct interface{}
}

func (v StructValidator) Validate() error {
	return validate.Struct(v.Struct)
}
