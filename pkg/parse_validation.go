package pkg

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidate(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(err.Field()),
				Message: "Field " + strings.ToLower(err.Field()) + " failed on " + err.Tag() + " rule",
			})
		}
	}
	return errors
}
