package providers

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/router"
	"CaliYa/cmd/api/router/groups"
	"CaliYa/config"
	"CaliYa/core/adapters/mongo"
	"CaliYa/core/adapters/postgres"
	"CaliYa/core/adapters/postgres/repo"
	"CaliYa/core/app"
	"CaliYa/core/utils"

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

	_ = Container.Provide(postgres.NewPostgresConnection)
	_ = Container.Provide(mongo.NewMongoConnection)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewProductsGroup)
	_ = Container.Provide(groups.NewOrdersGroup)
	_ = Container.Provide(groups.NewPromotionsGroup)
	_ = Container.Provide(groups.NewShopsGroup)
	_ = Container.Provide(groups.NewCustomersGroup)

	_ = Container.Provide(handler.NewProducts)
	_ = Container.Provide(handler.NewOrdersHandler)
	_ = Container.Provide(handler.NewPromos)
	_ = Container.Provide(handler.NewShopsHandler)
	_ = Container.Provide(handler.NewCustomersHandler)

	_ = Container.Provide(app.NewProductsApp)
	_ = Container.Provide(app.NewOrdersApp)
	_ = Container.Provide(app.NewPromotionsApp)
	_ = Container.Provide(app.NewShopsApp)
	_ = Container.Provide(app.NewCustomerApp)

	_ = Container.Provide(repo.NewProductsRepo)
	_ = Container.Provide(repo.NewOrdersRepo)
	_ = Container.Provide(repo.NewPromotionsRepository)
	_ = Container.Provide(repo.NewShopsRepository)
	_ = Container.Provide(repo.NewCustomersRepo)

	_ = Container.Provide(utils.NewHashPassword)

	return Container
}
