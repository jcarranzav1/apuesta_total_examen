package groups

import (
	"ApuestaTotal/internal/payment/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

type Payment interface {
	Resource(c *echo.Group)
}

type payment struct {
	paymentHandler handler.Payment
}

func NewPayment(productHandler handler.Payment) Payment {
	return &payment{
		productHandler,
	}
}

func (groups payment) Resource(route *echo.Group) {

	route.GET("/payment/:id", groups.paymentHandler.GetById)

}
