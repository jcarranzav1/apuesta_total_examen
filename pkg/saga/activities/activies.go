package activities

import (
	"context"
	"errors"
	"net/http"

	billDto "ApuestaTotal/internal/bill/domain/dto"
	entityBill "ApuestaTotal/internal/bill/domain/entity"
	portBill "ApuestaTotal/internal/bill/domain/ports"
	cartDto "ApuestaTotal/internal/cart/domain/dto"
	entityCart "ApuestaTotal/internal/cart/domain/entity"
	portCart "ApuestaTotal/internal/cart/domain/ports"
	paymentDto "ApuestaTotal/internal/payment/domain/dto"
	entityPayment "ApuestaTotal/internal/payment/domain/entity"
	portPayment "ApuestaTotal/internal/payment/domain/ports"
	productDto "ApuestaTotal/internal/products/domain/dto"
	portProducts "ApuestaTotal/internal/products/domain/ports"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Activities interface {
	GetProductByCart(ctx context.Context, cartId uint) (entityCart.Cart, error)
	GetStockByCart(ctx context.Context, products []entityCart.Item) ([]int, error)
	CreatePayment(ctx context.Context, paymentRequest paymentDto.PaymentCreate, products []entityCart.Item) (entityPayment.Payment, error)
	SuccessFullPayment(ctx context.Context, updateStatus paymentDto.UpdateStatus) error
	RejectedPayment(ctx context.Context, updateStatus paymentDto.UpdateStatus) error
	RemoveCart(ctx context.Context, cartId uint) (entityCart.Cart, error)
	RemovePayment(ctx context.Context, paymentId uint) error
	UpdateStock(ctx context.Context, products []entityCart.Item) error
	ReverseStock(ctx context.Context, products []entityCart.Item, baseQuantities []int) error
	CreateBill(ctx context.Context, paymentData entityPayment.Payment) (entityBill.Bill, error)
	RemoveBill(ctx context.Context, billId uint) error
}

type activities struct {
	billRepository     portBill.Bill
	paymentRepository  portPayment.Payment
	cartRepository     portCart.Cart
	productsRepository portProducts.Product
}

func NewActivitiesApplication(billRepository portBill.Bill,
	paymentRepository portPayment.Payment,
	cartRepository portCart.Cart,
	productsRepository portProducts.Product) Activities {
	return &activities{
		billRepository,
		paymentRepository,
		cartRepository,
		productsRepository,
	}
}

func (app *activities) GetProductByCart(ctx context.Context, cartId uint) (entityCart.Cart, error) {
	cart, err := app.cartRepository.GetById(ctx, cartId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entityCart.Cart{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return entityCart.Cart{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cart, nil
}

func (app *activities) GetStockByCart(ctx context.Context, products []entityCart.Item) ([]int, error) {
	var stock []int
	for _, product := range products {
		currentProduct, err := app.productsRepository.GetById(ctx, product.ProductID)
		if err != nil {
			return stock, echo.NewHTTPError(http.StatusInternalServerError, "Error to get the products")
		}
		stock = append(stock, currentProduct.Stock)
	}
	return stock, nil
}

func (app *activities) CreatePayment(ctx context.Context, paymentRequest paymentDto.PaymentCreate, products []entityCart.Item) (entityPayment.Payment, error) {

	for _, product := range products {
		currentProduct, err := app.productsRepository.GetById(ctx, product.ProductID)
		if err != nil {
			return entityPayment.Payment{}, echo.NewHTTPError(http.StatusInternalServerError, "Error to get the products")
		}
		paymentRequest.Products = append(paymentRequest.Products, entityPayment.Product{
			ID:       currentProduct.ID,
			Name:     currentProduct.Name,
			Price:    currentProduct.Price,
			Quantity: product.Quantity,
		})

		paymentRequest.Amount += currentProduct.Price * float64(product.Quantity)
	}

	//paymentRequest.Status = "pending"

	payment, err := app.paymentRepository.Create(ctx, paymentRequest)

	if err != nil {

		return entityPayment.Payment{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return payment, nil
}

func (app *activities) SuccessFullPayment(ctx context.Context, updateStatus paymentDto.UpdateStatus) error {

	err := app.paymentRepository.UpdateStatus(ctx, updateStatus)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *activities) RejectedPayment(ctx context.Context, updateStatus paymentDto.UpdateStatus) error {

	err := app.paymentRepository.UpdateStatus(ctx, updateStatus)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *activities) RemovePayment(ctx context.Context, paymentId uint) error {
	getPayment, _ := app.paymentRepository.GetById(ctx, paymentId)
	if getPayment.ID == 0 {
		return nil
	}

	err := app.paymentRepository.Remove(ctx, paymentId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *activities) RemoveCart(ctx context.Context, cartId uint) (entityCart.Cart, error) {

	cartRemoved, err := app.cartRepository.RemoveCart(ctx, cartId)

	if err != nil {
		return entityCart.Cart{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cartRemoved, nil
}

func (app *activities) RestoreCart(ctx context.Context, cartRestored entityCart.Cart) error {

	getCart, _ := app.cartRepository.GetById(ctx, cartRestored.ID)
	if getCart.ID != 0 {
		return nil
	}

	var createCart = cartDto.CreateProductDTO{
		ProductID: cartRestored.Items[0].ProductID,
		Quantity:  cartRestored.Items[0].Quantity,
	}

	newCart, err := app.cartRepository.CreateCart(ctx, createCart)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, product := range cartRestored.Items[1:] {
		var addProduct = cartDto.AddOrUpdateProductDTO{
			CartID:    newCart.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}

		err = app.cartRepository.AddProduct(ctx, addProduct)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return nil
}

func (app *activities) UpdateStock(ctx context.Context, products []entityCart.Item) error {
	for _, product := range products {
		currentProduct, err := app.productsRepository.GetById(ctx, product.ProductID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var productData productDto.ProductUpdate
		productData.Stock = currentProduct.Stock - product.Quantity
		productData.ID = product.ProductID

		err = app.productsRepository.Update(ctx, productData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}

func (app *activities) ReverseStock(ctx context.Context, products []entityCart.Item, baseQuantities []int) error {
	for index, product := range products {
		var productData productDto.ProductUpdate
		productData.Stock = baseQuantities[index]
		productData.ID = product.ProductID

		err := app.productsRepository.Update(ctx, productData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}

func (app *activities) CreateBill(ctx context.Context, paymentData entityPayment.Payment) (entityBill.Bill, error) {

	var createBill = billDto.BillCreate{
		PaymentMethod: paymentData.PaymentMethod,
		Currency:      paymentData.Currency,
		Amount:        paymentData.Amount,
	}
	for _, product := range paymentData.Products {
		createBill.Products = append(createBill.Products, entityBill.Product{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price,
		})

	}

	bill, err := app.billRepository.Create(ctx, createBill)

	if err != nil {
		return entityBill.Bill{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return bill, nil
}

func (app *activities) RemoveBill(ctx context.Context, billId uint) error {
	getBill, _ := app.billRepository.GetById(ctx, billId)
	if getBill.ID == 0 {
		return nil
	}

	err := app.billRepository.Remove(ctx, billId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
