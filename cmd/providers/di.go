package providers

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/handler/customers"
	"CaliYa/cmd/api/handler/order"
	middleware "CaliYa/cmd/api/middleware/requets"
	"CaliYa/cmd/api/router"
	"CaliYa/cmd/api/router/groups"
	"CaliYa/config"
	"CaliYa/core/adapters/mongo"
	"CaliYa/core/adapters/postgres"
	"CaliYa/core/adapters/postgres/repo"
	"CaliYa/core/adapters/postgres/repo/sessions"
	adapters "CaliYa/core/adapters/twilio"
	"CaliYa/core/app"
	"CaliYa/core/utils"

	sessionsHand "CaliYa/cmd/api/handler/sessions"

	customersRepo "CaliYa/core/adapters/postgres/repo/customers"
	customersApp "CaliYa/core/app/customers"
	"CaliYa/core/app/orders"
	sessionsApp "CaliYa/core/app/sessions"

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
	_ = Container.Provide(adapters.NewTwilioClient)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewProductsGroup)
	_ = Container.Provide(groups.NewOrdersGroup)
	_ = Container.Provide(groups.NewShopsGroup)
	_ = Container.Provide(groups.NewCustomersGroup)
	_ = Container.Provide(groups.NewSessionsGroup)

	_ = Container.Provide(middleware.NewRequestMiddleware)
	_ = Container.Provide(middleware.NewAuthMidlleware)

	_ = Container.Provide(handler.NewProducts)
	_ = Container.Provide(order.NewOrdersHandler)
	_ = Container.Provide(handler.NewShopsHandler)
	_ = Container.Provide(customers.NewCustomersHandler)
	_ = Container.Provide(customers.NewSignInCustomers)
	_ = Container.Provide(customers.NewCustomersAddressHandler)
	_ = Container.Provide(sessionsHand.NewSessionsHandler)

	_ = Container.Provide(sessionsApp.NewSessionsApp)
	_ = Container.Provide(app.NewProductsApp)
	_ = Container.Provide(orders.NewOrdersApp)
	_ = Container.Provide(app.NewShopsApp)
	_ = Container.Provide(customersApp.NewCustomerApp)
	_ = Container.Provide(customersApp.NewCustomersSessionsApp)
	_ = Container.Provide(customersApp.NewCustomersAddressApp)

	_ = Container.Provide(sessions.NewSessionsRepo)
	_ = Container.Provide(repo.NewProductsRepo)
	_ = Container.Provide(repo.NewOrdersRepo)
	_ = Container.Provide(repo.NewShopsRepository)
	_ = Container.Provide(customersRepo.NewCustomersRepo)
	_ = Container.Provide(customersRepo.NewCustomersAddressRepo)

	_ = Container.Provide(utils.NewHashPassword)
	_ = Container.Provide(utils.NewSessionUtils)

	return Container
}
