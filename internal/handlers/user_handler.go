package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
	"github.com/meedeley/go-launch-starter-code/pkg"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) {

}

func CreateUser(c *fiber.Ctx) error {
	var userReq entities.UserRegisterRequest

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	if err := c.BodyParser(userReq); err != nil {
		return pkg.ResponseJSON(c, 500, pkg.Response{
			Status:  500,
			Message: "Failed to parse request body",
			Data:    nil,
		}, nil)
	}

	db, _ := conf.NewPool()

	q := users.New(db)

	hashedPassword, err := hashPassword(userReq.Password)

	if err != nil {
		return err
	}

	q.CreateUser(ctx, users.CreateUserParams{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: hashedPassword,
	})

	return nil
}
