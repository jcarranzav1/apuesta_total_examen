package saga

import (
	"context"
	"net/http"

	entityBill "ApuestaTotal/internal/bill/domain/entity"
	"ApuestaTotal/internal/cart/domain/dto"
	entityCart "ApuestaTotal/internal/cart/domain/entity"
	entityPayment "ApuestaTotal/internal/payment/domain/entity"

	paymentDto "ApuestaTotal/internal/payment/domain/dto"
	"ApuestaTotal/pkg/saga/activities"
	paymentSagaDto "ApuestaTotal/pkg/saga/dto"
	"github.com/labstack/echo/v4"
)

type WorkFlow interface {
	PaymentSaga(context echo.Context) error
}

type paymentWorkflow struct {
	activitiesApp activities.Activities
	saga          Saga
}

func NewWorkFlow(activitiesApp activities.Activities, saga *Saga) WorkFlow {
	return &paymentWorkflow{
		activitiesApp,
		*saga,
	}
}

func (workflow *paymentWorkflow) PaymentSaga(ctxt echo.Context) error {
	ctx := ctxt.Request().Context()

	var sagaRequest paymentSagaDto.PaymentCreate
	var cart entityCart.Cart
	var stock []int
	var payment entityPayment.Payment
	var bill entityBill.Bill
	var createPayment = paymentDto.PaymentCreate{
		PaymentMethod: "Credit Card",
		Currency:      "Dollar",
	}

	if err := ctxt.Bind(&sagaRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		var err error
		cartId, _ := args[0].(uint)

		cart, err = workflow.activitiesApp.GetProductByCart(requestCtx, cartId)

		return err
	}, sagaRequest.CartId)

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		var err error
		products, _ := args[0].([]entityCart.Item)

		stock, err = workflow.activitiesApp.GetStockByCart(requestCtx, products)

		return err
	}, cart.Items)

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		var err error
		paymentRequest, _ := args[0].(paymentDto.PaymentCreate)
		products, _ := args[1].([]entityCart.Item)

		payment, err = workflow.activitiesApp.CreatePayment(requestCtx, paymentRequest, products)

		return err
	}, createPayment, cart.Items)

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		var err error
		cartId, _ := args[0].(uint)

		_, err = workflow.activitiesApp.RemoveCart(requestCtx, cartId)
		return err
	}, sagaRequest.CartId)

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		products, _ := args[0].([]entityCart.Item)
		return workflow.activitiesApp.UpdateStock(requestCtx, products)
	}, cart.Items)

	workflow.saga.AddOperation(func(requestCtx context.Context, args ...interface{}) error {
		paymentData, _ := args[0].(entityPayment.Payment)
		var err error
		bill, err = workflow.activitiesApp.CreateBill(requestCtx, paymentData)
		return err
	}, payment)

	// compensation
	workflow.saga.AddCompensation(func(requestCtx context.Context, args ...interface{}) error {
		var err error
		paymentId, _ := args[0].(uint)

		err = workflow.activitiesApp.RemovePayment(requestCtx, paymentId)
		return err
	}, payment.ID)

	workflow.saga.AddCompensation(func(requestCtx context.Context, args ...interface{}) error {
		var err error

		products, _ := args[0].([]entityCart.Item)
		baseStock, _ := args[1].([]int)

		err = workflow.activitiesApp.ReverseStock(requestCtx, products, baseStock)

		return err
	}, cart.Items, stock)

	workflow.saga.AddCompensation(func(requestCtx context.Context, args ...interface{}) error {
		billId, _ := args[0].(uint)

		var err error
		err = workflow.activitiesApp.RemoveBill(requestCtx, billId)
		return err
	}, bill.ID)

	err := workflow.saga.Execute(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctxt.JSON(http.StatusOK, dto.Message{
		Message: "Payment process completed successfully",
	})
}
