package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http/handlers"
	"go.uber.org/dig"
)

type RouteConfig struct {
	dig.In
	App                 *fiber.App
	UserHandler         *handlers.UserHandler
	ProtectedMiddleware fiber.Handler
	GuestMiddleware     fiber.Handler
}

func (c *RouteConfig) SetupRoutes() {
	c.App.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	c.setupGuestRoute()
	c.setupAuthRoute()
}

func (c *RouteConfig) setupGuestRoute() {
	c.App.Post("/register", c.UserHandler.Register, c.GuestMiddleware)
	c.App.Post("/login", c.UserHandler.Login, c.GuestMiddleware)
}

func (c *RouteConfig) setupAuthRoute() {
	c.App.Get("/users", c.UserHandler.FindAll, c.ProtectedMiddleware)
	c.App.Get("/user/:id", c.UserHandler.FindById, c.ProtectedMiddleware)
	c.App.Put("/user/:id", c.UserHandler.Update, c.ProtectedMiddleware)
	c.App.Delete("/user/:id", c.UserHandler.Delete, c.ProtectedMiddleware)
}
