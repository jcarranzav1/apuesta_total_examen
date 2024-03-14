package handler

import (
	"net/http"
	"strconv"

	"ApuestaTotal/internal/bill/application"
	"ApuestaTotal/internal/bill/domain/dto"
	"github.com/labstack/echo/v4"
)

const (
	getPaymentSuccess = "Get Product Success"
)

type Bill interface {
	GetById(context echo.Context) error
}

type billHandler struct {
	billApp application.Bill
}

func NewBillHandler(paymentApp application.Bill) Bill {
	return &billHandler{
		paymentApp,
	}
}

func (handler *billHandler) GetById(context echo.Context) error {
	ctx := context.Request().Context()

	billId, _ := strconv.Atoi(context.Param("id"))

	bill, err := handler.billApp.GetById(ctx, billId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: getPaymentSuccess,
		Data:    bill,
	})
}
