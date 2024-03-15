package dto

import "ApuestaTotal/internal/payment/domain/entity"

type PaymentCreate struct {
	Amount        float64          `json:"amount"`
	Currency      string           `json:"currency"`
	PaymentMethod string           `json:"payment_method"`
	Products      []entity.Product `json:"products"`
	Status        string           `json:"status"`
}

type UpdateStatus struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
