package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http/route"
	"github.com/meedeley/go-launch-starter-code/internal/provider"
)

func main() {
	container := provider.BuildContainer()

	err := container.Invoke(func(rc route.RouteConfig, app *fiber.App) {
		rc.SetupRoutes()

		app.Listen("localhost:3000")
	})

	if err != nil {
		log.Fatal("Dependency injection error:", err)
	}
}
