package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/handlers"
)

func Http() *fiber.App {

	app := conf.RunApp()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]any{
			"StatusCode": 200,
			"Message":    "Hello World",
		})
	})

	// => Default Prefix
	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Post("/register", handlers.Register)
	v1.Post("/login", handlers.Login)

	v1.Get("/findall", handlers.FindAllUser)

	return app
}
