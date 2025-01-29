package entities

import (
	"github.com/go-playground/validator/v10"
)

type UserRegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserRegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UserRegisterResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserLoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UserLoginResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
