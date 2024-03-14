package groups

import (
	"ApuestaTotal/internal/cart/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

type Cart interface {
	Resource(c *echo.Group)
}

type cart struct {
	cartHandler handler.Cart
}

func NewCart(cartHandler handler.Cart) Cart {
	return &cart{
		cartHandler,
	}
}

func (groups cart) Resource(route *echo.Group) {
	route.GET("/cart/:cartID", groups.cartHandler.GetById)
	route.POST("/cart", groups.cartHandler.CreateCart)
	route.POST("/cart/:cartID/product", groups.cartHandler.AddProduct)
	route.PUT("/cart/:cartID/product/:productID", groups.cartHandler.UpdateStock)
	route.DELETE("/cart/:cartID", groups.cartHandler.RemoveCart)
	route.DELETE("/cart/:cartID/product/:productID", groups.cartHandler.RemoveProduct)
}
