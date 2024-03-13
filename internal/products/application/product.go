package application

import (
	"context"
	"errors"
	"net/http"

	"ApuestaTotal/internal/products/domain/dto"
	"ApuestaTotal/internal/products/domain/entity"
	"ApuestaTotal/internal/products/domain/ports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Product interface {
	Create(ctx context.Context, newProduct dto.ProductCreate) (entity.Product, error)
	Update(ctx context.Context, updateProduct dto.ProductUpdate) error
	GetById(ctx context.Context, id int) (entity.Product, error)
	GetAll(ctx context.Context) ([]entity.Product, error)
}

type productApp struct {
	productRepository ports.Product
}

func NewProductApplication(alertsRepo ports.Product) Product {
	return &productApp{
		alertsRepo,
	}
}

func (app *productApp) GetById(ctx context.Context, id int) (entity.Product, error) {
	product, err := app.productRepository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return entity.Product{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return product, nil
}

func (app *productApp) GetAll(ctx context.Context) ([]entity.Product, error) {
	products, err := app.productRepository.GetAll(ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Product{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return []entity.Product{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return products, nil
}

func (app *productApp) Create(ctx context.Context, newProduct dto.ProductCreate) (entity.Product, error) {
	productCreated, err := app.productRepository.Create(ctx, newProduct)

	if err != nil {
		return entity.Product{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return productCreated, nil
}

func (app *productApp) Update(ctx context.Context, updateProduct dto.ProductUpdate) error {
	_, err := app.productRepository.GetById(ctx, updateProduct.ID)

	if err != nil {
		return err
	}

	err = app.productRepository.Update(ctx, updateProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
