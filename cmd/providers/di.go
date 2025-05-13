package providers

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/router"
	"CaliYa/cmd/api/router/groups"
	"CaliYa/config"
	"CaliYa/core/adapters"
	"CaliYa/core/adapters/repo"

	"CaliYa/core/app"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() config.Config {
		config.Environments()
		return *config.Get()
	})

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(adapters.NewPostgresConnection)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewProductsGroup)

	_ = Container.Provide(handler.NewProducts)

	_ = Container.Provide(app.NewProductsApp)

	_ = Container.Provide(repo.NewProductsRepo)

	return Container
}
