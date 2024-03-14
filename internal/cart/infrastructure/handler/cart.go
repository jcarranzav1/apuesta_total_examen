package handler

import (
	"net/http"
	"strconv"

	"ApuestaTotal/internal/cart/application"
	"ApuestaTotal/internal/cart/domain/dto"
	"github.com/labstack/echo/v4"
)

const (
	getCartSuccess       = "Get Cart Success"
	createCartSuccess    = "Cart Created Successfully"
	removeCartSuccess    = "Cart Removed Successfully"
	addProductSuccess    = "Product Added Successfully"
	removeProductSuccess = "Product Removed Successfully"
	updateStockSuccess   = "Stock Updated Successfully"
)

type Cart interface {
	GetById(context echo.Context) error
	CreateCart(context echo.Context) error
	RemoveCart(context echo.Context) error
	AddProduct(context echo.Context) error
	RemoveProduct(context echo.Context) error
	UpdateStock(context echo.Context) error
}

type cartHandler struct {
	cartApp application.Cart
}

func NewCartHandler(cartApp application.Cart) Cart {
	return &cartHandler{
		cartApp,
	}
}

func (handler *cartHandler) GetById(context echo.Context) error {
	ctx := context.Request().Context()

	cartId, _ := strconv.Atoi(context.Param("cartID"))

	cart, err := handler.cartApp.GetById(ctx, uint(cartId))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: getCartSuccess,
		Data:    cart,
	})
}

func (handler *cartHandler) CreateCart(context echo.Context) error {

	ctx := context.Request().Context()

	var newCart dto.CreateProductDTO

	if err := context.Bind(&newCart); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := newCart.ValidateCreate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cartCreated, err := handler.cartApp.CreateCart(ctx, newCart)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: createCartSuccess,
		Data:    cartCreated,
	})
}

func (handler *cartHandler) RemoveCart(context echo.Context) error {

	ctx := context.Request().Context()

	cartId, _ := strconv.Atoi(context.Param("cartID"))

	err := handler.cartApp.RemoveCart(ctx, uint(cartId))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: removeCartSuccess,
	})
}

func (handler *cartHandler) AddProduct(context echo.Context) error {

	ctx := context.Request().Context()

	cartID, _ := strconv.Atoi(context.Param("cartID"))

	var addProduct dto.AddOrUpdateProductDTO
	addProduct.CartID = uint(cartID)

	if err := context.Bind(&addProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := addProduct.ValidateAddOrUpdate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := handler.cartApp.AddProduct(ctx, addProduct)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: addProductSuccess,
	})
}

func (handler *cartHandler) RemoveProduct(context echo.Context) error {

	ctx := context.Request().Context()

	cartID, _ := strconv.Atoi(context.Param("cartID"))
	productID, _ := strconv.Atoi(context.Param("productID"))

	var removeProduct dto.RemoveProductDTO
	removeProduct.CartID = uint(cartID)
	removeProduct.ProductID = uint(productID)

	if err := context.Bind(&removeProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := removeProduct.ValidateRemove(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := handler.cartApp.RemoveProduct(ctx, removeProduct)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: removeProductSuccess,
	})
}

func (handler *cartHandler) UpdateStock(context echo.Context) error {
	ctx := context.Request().Context()

	cartID, _ := strconv.Atoi(context.Param("cartID"))
	productID, _ := strconv.Atoi(context.Param("productID"))

	var cart dto.AddOrUpdateProductDTO
	cart.CartID = uint(cartID)
	cart.ProductID = uint(productID)

	if err := context.Bind(&cart); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := cart.ValidateAddOrUpdate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := handler.cartApp.UpdateStock(ctx, cart)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: updateStockSuccess,
	})
}
