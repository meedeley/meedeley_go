package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/handlers"
)

func Http() *fiber.App {

	app := conf.RunApp()

	// => Default Prefix
	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Post("/register", handlers.CreateUser)

	return app
}
