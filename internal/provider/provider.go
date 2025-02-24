package provider

import (
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http/handlers"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http/middlewares"
	"github.com/meedeley/go-launch-starter-code/internal/usecase"
	"go.uber.org/dig"
)

func registerApp(container *dig.Container) {
	container.Provide(conf.ProvideApp)
	container.Provide(conf.NewPool)
}

func registerMiddleware(container *dig.Container) {
	container.Provide(middlewares.NewAuthMiddleware().GuestOnly)
	container.Provide(middlewares.NewAuthMiddleware().Protected)
}

func registerHandler(container *dig.Container) {
	container.Provide(handlers.NewUserHandler)
}

func registerUseCase(container *dig.Container) {
	container.Provide(usecase.NewUserUseCase)
}

func registerRespository(container *dig.Container) {
	// container.Provide()
}

func BuildProvider() *dig.Container {
	container := dig.New()

	registerApp(container)
	registerRespository(container)
	registerUseCase(container)
	registerHandler(container)
	registerMiddleware(container)

	return container
}
