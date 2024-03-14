package repository

import (
	"context"

	"ApuestaTotal/internal/payment/domain/dto"
	"ApuestaTotal/internal/payment/domain/entity"
	"ApuestaTotal/internal/payment/domain/ports"
	"ApuestaTotal/internal/payment/infrastructure/adapters/model"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) ports.Payment {
	return &paymentRepository{
		db,
	}
}

func (repository paymentRepository) GetById(ctx context.Context, id int) (entity.Payment, error) {
	var modelPayment model.Payment

	if result := repository.db.WithContext(ctx).First(&modelPayment, id); result.Error != nil {
		return entity.Payment{}, result.Error
	}

	return modelPayment.ToPaymentDomain(), nil
}

func (repository paymentRepository) Create(ctx context.Context, newPayment dto.PaymentCreate) (entity.Payment, error) {

	var modelProducts []model.Product
	for _, item := range newPayment.Products {
		modelProducts = append(modelProducts, model.Product{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	var modelPayment = model.Payment{
		Amount:        newPayment.Amount,
		Currency:      newPayment.Currency,
		PaymentMethod: newPayment.PaymentMethod,
		Products:      modelProducts,
	}

	if err := repository.db.WithContext(ctx).
		Create(&modelPayment).
		Error; err != nil {

		return entity.Payment{}, err
	}
	return modelPayment.ToPaymentDomain(), nil
}
