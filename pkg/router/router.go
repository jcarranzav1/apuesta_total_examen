package router

import (
	routeBill "ApuestaTotal/internal/bill/infrastructure/groups"
	routeCart "ApuestaTotal/internal/cart/infrastructure/groups"
	routePayment "ApuestaTotal/internal/payment/infrastructure/groups"
	routeProducts "ApuestaTotal/internal/products/infrastructure/groups"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server  *echo.Echo
	product routeProducts.Product
	cart    routeCart.Cart
	payment routePayment.Payment
	bill    routeBill.Bill
}

func New(
	server *echo.Echo,
	product routeProducts.Product,
	cart routeCart.Cart,
	payment routePayment.Payment,
	bill routeBill.Bill,

) *Router {
	return &Router{
		server,
		product,
		cart,
		payment,
		bill,
	}
}

func (router *Router) Init() {
	router.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	router.server.Use(middleware.Recover())

	basePath := router.server.Group("/api")
	basePath.GET("/health", HealthCheck)

	//if config.IsDevEnvironment() {
	//	basePath.GET("/swagger/*", echoSwagger.WrapHandler)
	//}

	router.product.Resource(basePath)
	router.cart.Resource(basePath)
	router.payment.Resource(basePath)
	router.bill.Resource(basePath)

}
