package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email,unique_email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserRegisterRequest) Validate() []pkg.ValidationError {
	validate := validator.New()
	validate.RegisterValidation("unique_email", pkg.UniqueEmailValidator)

	err := validate.Struct(r)

	return pkg.ParseValidate(err)
}

type UserRegisterResponse struct {
	Id        any        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserLoginRequest) Validate() []pkg.ValidationError {
	validate := validator.New()

	return pkg.ParseValidate(validate.Struct(r))
}

type UserLoginResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email,min=2"`
}

func (r *UpdateUserRequest) Validate() []pkg.ValidationError {
	validate := validator.New()
	err := validate.Struct(r)

	return pkg.ParseValidate(err)
}

type UpdateUserResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
