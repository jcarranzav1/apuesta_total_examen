package handler

import (
	"net/http"
	"strconv"

	"ApuestaTotal/internal/products/application"
	"ApuestaTotal/internal/products/domain/dto"
	"github.com/labstack/echo/v4"
)

const (
	getProductSuccess    = "Get Product Success"
	getAllProductSuccess = "Get All Items Success"
	createProductSuccess = "Product Created Successfully"
	updateProductSuccess = "Product Updated Successfully"
)

type Product interface {
	GetById(context echo.Context) error
	GetAll(context echo.Context) error
	Create(context echo.Context) error
	Update(context echo.Context) error
}

type productHandler struct {
	productApp application.Product
}

func NewProductHandler(productApp application.Product) Product {
	return &productHandler{
		productApp,
	}
}

func (handler *productHandler) GetById(context echo.Context) error {
	ctx := context.Request().Context()

	productId, _ := strconv.Atoi(context.Param("id"))

	product, err := handler.productApp.GetById(ctx, productId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: getProductSuccess,
		Data:    product,
	})
}

func (handler *productHandler) GetAll(context echo.Context) error {
	ctx := context.Request().Context()

	products, err := handler.productApp.GetAll(ctx)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: getAllProductSuccess,
		Data:    products,
	})
}

func (handler *productHandler) Create(context echo.Context) error {

	ctx := context.Request().Context()

	var newProduct dto.ProductCreate

	if err := context.Bind(&newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := newProduct.ValidateCreate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	productCreated, err := handler.productApp.Create(ctx, newProduct)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: createProductSuccess,
		Data:    productCreated,
	})
}

func (handler *productHandler) Update(context echo.Context) error {
	ctx := context.Request().Context()

	productId, _ := strconv.Atoi(context.Param("id"))

	var product dto.ProductUpdate

	product.ID = productId

	if err := context.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := product.ValidateUpdate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := handler.productApp.Update(ctx, product)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: updateProductSuccess,
	})
}
