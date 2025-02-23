package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

type User struct {
	Id        int        `json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password,omitempty" form:"password"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (r *UserRegisterRequest) Validate() []pkg.ValidationError {
	validate := validator.New()
	err := validate.Struct(r)

	return pkg.ParseValidate(err)
}

type UserRegisterResponse struct {
	Id    any    `json:"id"`
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
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
