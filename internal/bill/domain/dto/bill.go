package dto

import "ApuestaTotal/internal/bill/domain/entity"

type BillCreate struct {
	PaymentID     uint             `json:"payment_id"`
	Amount        float64          `json:"amount"`
	Currency      string           `json:"currency"`
	PaymentMethod string           `json:"payment_method"`
	Products      []entity.Product `json:"products"`
}
