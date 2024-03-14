package handler

import (
	"net/http"
	"strconv"

	"ApuestaTotal/internal/products/application"
	"ApuestaTotal/internal/products/domain/dto"
	"github.com/labstack/echo/v4"
)

const (
	getPaymentSuccess = "Get Product Success"
)

type Payment interface {
	GetById(context echo.Context) error
}

type paymentHandler struct {
	paymentApp application.Product
}

func NewPaymentHandler(paymentApp application.Product) Payment {
	return &paymentHandler{
		paymentApp,
	}
}

func (handler *paymentHandler) GetById(context echo.Context) error {
	ctx := context.Request().Context()

	paymentId, _ := strconv.Atoi(context.Param("id"))

	payment, err := handler.paymentApp.GetById(ctx, paymentId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: getPaymentSuccess,
		Data:    payment,
	})
}
