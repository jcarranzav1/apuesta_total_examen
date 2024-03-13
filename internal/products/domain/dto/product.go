package dto

import "github.com/go-playground/validator/v10"

type ProductCreate struct {
	Name  string  `json:"name" validate:"required,min=3,max=60"`
	Price float64 `json:"price" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
}

type ProductUpdate struct {
	ID    int     `json:"id" validate:"required"`
	Name  string  `json:"name" validate:"max=60"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (alert *ProductCreate) ValidateCreate() error {
	validate := validator.New()
	if err := validate.Struct(alert); err != nil {
		return err
	}

	return nil
}

func (alert *ProductUpdate) ValidateUpdate() error {
	validate := validator.New()

	if err := validate.Struct(alert); err != nil {
		return err
	}

	return nil

}
