package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http/route"
	"github.com/meedeley/go-launch-starter-code/internal/provider"
)

func main() {
	db, _ := conf.NewPool()
	container := provider.BuildProvider()

	err := container.Invoke(func(rc route.RouteConfig, app *fiber.App) {
		rc.SetupRoutes()

		app.Listen("localhost:3000")
	})

	defer db.Close()

	if err != nil {
		log.Fatal("Dependency injection error:", err)
	}
}
