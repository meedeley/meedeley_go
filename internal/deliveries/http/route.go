package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

func Http() *fiber.App {

	app := conf.RunAppWithGracefulShutdown()

	// => Default Prefix
	api := app.Group("api")

	v1 := api.Group("v1")
	v1.Get("/data", func(c *fiber.Ctx) error {
		return pkg.ResponseJSON(c, 200, entities.User{
			Id:       1,
			Name:     "Nichola Saputra",
			Email:    "nicholasaputra77@gmail.com",
			Password: "!Mebelopik123",
		},
			nil)
	})

	return app
}
