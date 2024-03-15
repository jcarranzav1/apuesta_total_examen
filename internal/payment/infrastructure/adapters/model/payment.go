package model

import (
	"ApuestaTotal/internal/payment/domain/entity"
	"gorm.io/gorm"
)

type Product struct {
	ID       uint
	Name     string
	Price    float64
	Quantity int
}

type Payment struct {
	gorm.Model
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	PaymentMethod string    `json:"payment_method"`
	Products      []Product `gorm:"type:json" json:"products"`
	Status        string    `json:"´status"`
}

func (p *Payment) ToPaymentDomain() entity.Payment {
	var domainProducts []entity.Product
	for _, item := range p.Products {
		domainProducts = append(domainProducts, entity.Product{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}
	return entity.Payment{
		ID:            p.ID,
		Amount:        p.Amount,
		Currency:      p.Currency,
		PaymentMethod: p.PaymentMethod,
		Products:      domainProducts,
		Status:        p.Status,
	}
}
