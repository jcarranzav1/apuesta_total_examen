package dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateProductDTO struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
}

type AddOrUpdateProductDTO struct {
	Quantity  int  `json:"quantity" validate:"required"`
	ProductID uint `json:"product_id" validate:"required"`
	CartID    uint `json:"cart_id"`
}

type RemoveProductDTO struct {
	ProductID uint `json:"product_id"`
	CartID    uint `json:"cart_id"`
}

func (alert *CreateProductDTO) ValidateCreate() error {
	validate := validator.New()
	if err := validate.Struct(alert); err != nil {
		return err
	}

	return nil
}

func (alert *AddOrUpdateProductDTO) ValidateAddOrUpdate() error {
	validate := validator.New()
	if err := validate.Struct(alert); err != nil {
		return err
	}

	return nil
}

func (alert *RemoveProductDTO) ValidateRemove() error {
	validate := validator.New()
	if err := validate.Struct(alert); err != nil {
		return err
	}

	return nil
}
