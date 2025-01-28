package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Http() error {

	app := fiber.New(fiber.Config{
		AppName:       "Meedeley Go Starter",
		StrictRouting: true,
		ErrorHandler:  fiber.DefaultErrorHandler,
		IdleTimeout:   time.Hour * 1,
	})

	v1 := app.Group("v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app.Listen(":3000")
}
