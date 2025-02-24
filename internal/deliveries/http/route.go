package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/deliveries/middleware"
	"github.com/meedeley/go-launch-starter-code/internal/handlers"
)

func Http() *fiber.App {

	app := conf.RunApp()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	// => Default Prefix
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Post("/register", handlers.Register, middleware.GuestOnly())
	v1.Post("/login", handlers.Login, middleware.GuestOnly())

	// => CRUD USER OPERATION
	v1.Get("/users", handlers.FindAllUser, middleware.Protected())
	v1.Get("/user/:id", handlers.FindUserById, middleware.Protected())
	v1.Put("/user/:id", handlers.UpdateUser, middleware.Protected())
	v1.Delete("/user/:id", handlers.DeleteUser, middleware.Protected())

	// => CRUD EXAMPLE OPERATION
	v1.Get("/example", handlers.FindAllExample)

	return app
}
