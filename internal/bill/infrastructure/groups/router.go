package groups

import (
	"ApuestaTotal/internal/bill/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

type Bill interface {
	Resource(c *echo.Group)
}

type bill struct {
	billHandler handler.Bill
}

func NewBill(productHandler handler.Bill) Bill {
	return &bill{
		productHandler,
	}
}

func (groups bill) Resource(route *echo.Group) {

	route.GET("/bill/:id", groups.billHandler.GetById)

}
