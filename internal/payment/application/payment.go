package application

import (
	"context"
	"errors"
	"net/http"

	"ApuestaTotal/internal/payment/domain/entity"
	"ApuestaTotal/internal/payment/domain/ports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Payment interface {
	GetById(ctx context.Context, id uint) (entity.Payment, error)
}

type paymentApp struct {
	paymentRepository ports.Payment
}

func NewPaymentApplication(paymentRepo ports.Payment) Payment {
	return &paymentApp{
		paymentRepo,
	}
}

func (app *paymentApp) GetById(ctx context.Context, id uint) (entity.Payment, error) {
	payment, err := app.paymentRepository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Payment{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return entity.Payment{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return payment, nil
}
