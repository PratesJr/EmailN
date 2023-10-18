package validators

import (
	"emailn/internal/exceptins"
	"errors"
	"github.com/go-playground/validator/v10"
)

func Validate(object interface{}) error {
	validate := validator.New()
	err := validate.Struct(object)
	if nil == err {
		return nil
	}
	validationErr := err.(validator.ValidationErrors)[0]

	switch validationErr.Tag() {
	case "required":
		return errors.New(validationErr.StructField() + " is required")
	case "min":
		return errors.New(validationErr.StructField() + " is required with min" + validationErr.Param())
	case "max":
		return errors.New(validationErr.StructField() + " is required with max" + validationErr.Param())
	case "email":
		return errors.New(validationErr.StructField() + " is not valid")
	}

	return exceptins.UnkownErrror
}
