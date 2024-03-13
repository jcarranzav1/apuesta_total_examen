package router

import (
	"ApuestaTotal/internal/products/infrastructure/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server         *echo.Echo
	productHandler handler.Product
}

func New(
	server *echo.Echo,
	productHandler handler.Product,
) *Router {
	return &Router{
		server,
		productHandler,
	}
}

func (router *Router) Init() {
	router.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	router.server.Use(middleware.Recover())

	basePath := router.server.Group("/api")
	basePath.GET("/health", handler.HealthCheck)

	//if config.IsDevEnvironment() {
	//	basePath.GET("/swagger/*", echoSwagger.WrapHandler)
	//}

	basePath.POST("/products", router.productHandler.Create)
	basePath.GET("/products", router.productHandler.GetAll)
	basePath.GET("/products/:id", router.productHandler.GetById)
	basePath.PUT("/products/:id", router.productHandler.Update)

}
