package formatter

import "github.com/go-playground/validator/v10"

func MapValidationErr(err error) map[string]string {
	validationErrs := err.(validator.ValidationErrors)
	errors := map[string]string{}
	for _, e := range validationErrs {
		errors[e.Field()] = e.Tag()
	}
	return errors
}
