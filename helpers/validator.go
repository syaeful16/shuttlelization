package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(v interface{}) map[string]string {
	validate := validator.New()
	err := validate.Struct(v)

	if err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			// Ubah huruf pertama field menjadi huruf kecil
			fieldName := strings.ToLower(e.Field()[:1]) + e.Field()[1:]

			// Custom error messages
			switch e.Tag() {
			case "required":
				errors[fieldName] = e.Field() + " is required"
			case "min":
				errors[fieldName] = e.Field() + " must be at least " + e.Param() + " characters"
			case "max":
				errors[fieldName] = e.Field() + " must be at most " + e.Param() + " characters"
			case "email":
				errors[fieldName] = "Invalid email format"
			case "numeric":
				errors[fieldName] = e.Field() + " must be a number"
			default:
				errors[fieldName] = "Invalid " + e.Field()
			}
		}
		return errors
	}
	return nil
}
