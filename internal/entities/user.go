package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"string"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
