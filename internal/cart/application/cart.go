package application

import (
	"context"
	"errors"
	"net/http"

	"ApuestaTotal/internal/cart/domain/dto"
	"ApuestaTotal/internal/cart/domain/entity"
	"ApuestaTotal/internal/cart/domain/ports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Cart interface {
	GetById(ctx context.Context, id uint) (entity.Cart, error)
	CreateCart(ctx context.Context, newCart dto.CreateProductDTO) (entity.Cart, error)
	RemoveCart(ctx context.Context, id uint) (entity.Cart, error)
	AddProduct(ctx context.Context, newProduct dto.AddOrUpdateProductDTO) error
	RemoveProduct(ctx context.Context, removeProduct dto.RemoveProductDTO) error
	UpdateStock(ctx context.Context, updateProduct dto.AddOrUpdateProductDTO) error
}

type cartApp struct {
	cartRepository ports.Cart
}

func NewCartApplication(cartRepo ports.Cart) Cart {
	return &cartApp{
		cartRepo,
	}
}

func (app *cartApp) GetById(ctx context.Context, id uint) (entity.Cart, error) {
	cart, err := app.cartRepository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Cart{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return entity.Cart{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cart, nil
}

func (app *cartApp) CreateCart(ctx context.Context, newCart dto.CreateProductDTO) (entity.Cart, error) {
	cartCreated, err := app.cartRepository.CreateCart(ctx, newCart)

	if err != nil {
		return entity.Cart{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cartCreated, nil
}

func (app *cartApp) RemoveCart(ctx context.Context, id uint) (entity.Cart, error) {
	cartRemoved, err := app.cartRepository.RemoveCart(ctx, id)

	if err != nil {
		return entity.Cart{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cartRemoved, nil
}

func (app *cartApp) AddProduct(ctx context.Context, newProduct dto.AddOrUpdateProductDTO) error {
	err := app.cartRepository.AddProduct(ctx, newProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *cartApp) RemoveProduct(ctx context.Context, removeProduct dto.RemoveProductDTO) error {
	err := app.cartRepository.RemoveProduct(ctx, removeProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *cartApp) UpdateStock(ctx context.Context, updateProduct dto.AddOrUpdateProductDTO) error {
	err := app.cartRepository.UpdateProduct(ctx, updateProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
