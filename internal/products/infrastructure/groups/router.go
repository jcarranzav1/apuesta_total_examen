package groups

import (
	"ApuestaTotal/internal/products/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

type Product interface {
	Resource(c *echo.Group)
}

type product struct {
	productHandler handler.Product
}

func NewProduct(productHandler handler.Product) Product {
	return &product{
		productHandler,
	}
}

func (groups product) Resource(route *echo.Group) {

	route.GET("/products", groups.productHandler.GetAll)
	route.GET("/products/:id", groups.productHandler.GetById)
	route.POST("/products", groups.productHandler.Create)
	route.PUT("/products/:id", groups.productHandler.Update)
}
