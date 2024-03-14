package providers

import (
	appBill "ApuestaTotal/internal/bill/application"
	repoBill "ApuestaTotal/internal/bill/infrastructure/adapters/repository"
	groupBill "ApuestaTotal/internal/bill/infrastructure/groups"
	handlerBill "ApuestaTotal/internal/bill/infrastructure/handler"
	appCart "ApuestaTotal/internal/cart/application"
	repoCart "ApuestaTotal/internal/cart/infrastructure/adapters/repository"
	groupCart "ApuestaTotal/internal/cart/infrastructure/groups"
	handlerCart "ApuestaTotal/internal/cart/infrastructure/handler"
	appPayment "ApuestaTotal/internal/payment/application"
	repoPayment "ApuestaTotal/internal/payment/infrastructure/adapters/repository"
	groupPayment "ApuestaTotal/internal/payment/infrastructure/groups"
	handlerPayment "ApuestaTotal/internal/payment/infrastructure/handler"
	appProduct "ApuestaTotal/internal/products/application"
	repoProduct "ApuestaTotal/internal/products/infrastructure/adapters/repository"
	groupProduct "ApuestaTotal/internal/products/infrastructure/groups"
	handlerProduct "ApuestaTotal/internal/products/infrastructure/handler"
	"ApuestaTotal/pkg/database"
	"ApuestaTotal/pkg/router"
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

	_ = Container.Provide(repoProduct.NewProductRepository)
	_ = Container.Provide(repoCart.NewCartRepository)
	_ = Container.Provide(repoPayment.NewPaymentRepository)
	_ = Container.Provide(repoBill.NewBillRepository)

	_ = Container.Provide(appProduct.NewProductApplication)
	_ = Container.Provide(appCart.NewCartApplication)
	_ = Container.Provide(appPayment.NewPaymentApplication)
	_ = Container.Provide(appBill.NewBillApplication)

	_ = Container.Provide(handlerProduct.NewProductHandler)
	_ = Container.Provide(handlerCart.NewCartHandler)
	_ = Container.Provide(handlerPayment.NewPaymentHandler)
	_ = Container.Provide(handlerBill.NewBillHandler)

	_ = Container.Provide(groupProduct.NewProduct)
	_ = Container.Provide(groupCart.NewCart)
	_ = Container.Provide(groupPayment.NewPayment)
	_ = Container.Provide(groupBill.NewBill)

	_ = Container.Provide(router.New)

	return Container
}
