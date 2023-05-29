package helper

import "github.com/go-playground/validator/v10"

func RegisterValidation(validation *validator.Validate) *validator.Validate {
	return validation
}
