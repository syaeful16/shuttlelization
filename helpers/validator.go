package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/syaeful16/shuttlelization/database"
	"gorm.io/gorm"
)

// Validator untuk memastikan nilai unik dalam tabel yang berbeda
func UniqueValidator(db *gorm.DB) validator.Func {
	return func(fl validator.FieldLevel) bool {
		var count int64
		value := fl.Field().String()

		// Ambil parameter dari tag validasi, contoh: "customer_logins.email"
		params := strings.Split(fl.Param(), ".")
		if len(params) != 2 {
			return false
		}

		table := params[0]
		column := params[1]

		// Query database
		err := db.Table(table).Where(column+" = ?", value).Count(&count).Error
		if err != nil {
			return false
		}

		return count == 0 // True jika data belum ada di database
	}
}

func Validate(v interface{}) map[string]string {
	validate := validator.New()

	// Registrasi validator unik (untuk berbagai tabel & kolom)
	validate.RegisterValidation("unique", UniqueValidator(database.DB))

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
			case "unique":
				errors[fieldName] = e.Field() + " already exists"
			default:
				errors[fieldName] = "Invalid " + e.Field()
			}
		}
		return errors
	}
	return nil
}
