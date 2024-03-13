package providers

import (
	"ApuestaTotal/internal/products/application"
	"ApuestaTotal/internal/products/infrastructure/adapters/repository"
	"ApuestaTotal/internal/products/infrastructure/handler"
	"ApuestaTotal/internal/products/infrastructure/router"
	"ApuestaTotal/pkg/database"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(database.NewConnection)

	_ = Container.Provide(repository.NewAudienceRepository)
	_ = Container.Provide(application.NewProductApplication)
	_ = Container.Provide(handler.NewProductHandler)
	_ = Container.Provide(router.New)

	return Container
}
