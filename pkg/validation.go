package pkg

import (
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidate(err error) []ValidationError {
	validated, _ := err.(validator.ValidationErrors)

	var errors []ValidationError
	for _, e := range validated {
		var message string

		// => Registered Custom Validation At Here
		switch e.Tag() {
		case "unique_email":
			message = "email already exist"
		default:
			message = "field " + strings.ToLower(e.Field()) + " " + e.Tag()
		}

		errors = append(errors, ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: message,
		})
	}

	return errors
}

func UniqueEmailValidator(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	db, _ := conf.NewPool()

	q := users.New(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := q.FindUserByEmail(ctx, email)
	return err != nil
}
