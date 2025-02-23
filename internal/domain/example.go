package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Example struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleRequest struct {
	Id          int    `json:"id" form:"id" validate:"required,numeric"`
	Name        string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Image       string `json:"image" fom:"image" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

func (r *ExampleRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type ExampleResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
