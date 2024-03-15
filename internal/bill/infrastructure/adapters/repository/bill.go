package repository

import (
	"context"

	"ApuestaTotal/internal/bill/domain/dto"
	"ApuestaTotal/internal/bill/domain/entity"
	"ApuestaTotal/internal/bill/domain/ports"
	"ApuestaTotal/internal/bill/infrastructure/adapters/model"
	"gorm.io/gorm"
)

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) ports.Bill {
	return &billRepository{
		db,
	}
}

func (repository billRepository) GetById(ctx context.Context, id uint) (entity.Bill, error) {
	var modelBill = model.Bill{}

	if result := repository.db.WithContext(ctx).First(&modelBill, id); result.Error != nil {
		return entity.Bill{}, result.Error
	}

	return modelBill.ToBillDomain(), nil
}

func (repository billRepository) Create(ctx context.Context, newPayment dto.BillCreate) (entity.Bill, error) {

	var modelProducts []model.Product
	for _, item := range newPayment.Products {
		modelProducts = append(modelProducts, model.Product{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	var modelBill = model.Bill{
		Amount:        newPayment.Amount,
		Currency:      newPayment.Currency,
		PaymentMethod: newPayment.PaymentMethod,
		Products:      modelProducts,
	}

	if err := repository.db.WithContext(ctx).
		Create(&modelBill).
		Error; err != nil {

		return entity.Bill{}, err
	}
	return modelBill.ToBillDomain(), nil
}

func (repository billRepository) Remove(ctx context.Context, id uint) error {
	var modelBill = model.Bill{}

	if err := repository.db.WithContext(ctx).
		Unscoped().
		Delete(&modelBill, id).Error; err != nil {
		return err
	}
	return nil
}
