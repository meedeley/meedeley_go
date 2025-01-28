package entities

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"string"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
