package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

func Http() error {

	app := conf.RunAppWithGracefulShutdown()

	v1 := app.Group("v1")
	v1.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(3 * time.Second)
		return pkg.ResponseJSON(c, 200, entities.User{
			Id:       1,
			Name:     "Nichola Saputra",
			Email:    "nicholasaputra77@gmail.com",
			Password: "!Mebelopik123",
		}, nil)
	})

	return app.Listen(":3000")
}
