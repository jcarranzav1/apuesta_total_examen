package dto

type PaymentCreate struct {
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
	CartId        uint   `json:"cart-id"`
}
