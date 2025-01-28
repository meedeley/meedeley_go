package handlers

import "github.com/meedeley/go-launch-starter-code/internal/entities"

type UserHandler interface {
	Register() (entities.User, error)
	Login() (entities.User, error)
	Logout() (entities.User, error)
}

func Register() {

}

func Login() {

}

func Logout() {

}
