package application

import (
	"context"
	"errors"
	"net/http"

	"ApuestaTotal/internal/bill/domain/entity"
	"ApuestaTotal/internal/bill/domain/ports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Bill interface {
	GetById(ctx context.Context, id int) (entity.Bill, error)
}

type billApp struct {
	billRepository ports.Bill
}

func NewBillApplication(paymentRepo ports.Bill) Bill {
	return &billApp{
		paymentRepo,
	}
}

func (app *billApp) GetById(ctx context.Context, id int) (entity.Bill, error) {
	bill, err := app.billRepository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Bill{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return entity.Bill{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return bill, nil
}
