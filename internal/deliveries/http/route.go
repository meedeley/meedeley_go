package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/handlers"
)

func Http() *fiber.App {

	app := conf.RunApp()

	// => Default Prefix
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Post("/register", handlers.Register)
	v1.Post("/login", handlers.Login)

	// => CRUD USER OPERATION
	v1.Get("/users", handlers.FindAllUser)
	v1.Get("/user/:id", handlers.FindUserById)
	v1.Put("/user/:id", handlers.UpdateUser)
	v1.Delete("/user/:id", handlers.DeleteUser)

	// => CRUD EXAMPLE OPERATION

	return app
}
