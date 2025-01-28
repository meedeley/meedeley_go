package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
)

func RegisterHandler(c *fiber.Ctx) error {
	user := new(entities.User)

	err := c.BodyParser(user)

	if err != nil {
		return err
	}

	


}

func LoginHandler() {

}

func LogoutHandler() {

}
