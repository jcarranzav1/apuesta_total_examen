package entity

import "time"

type Product struct {
	ID       uint
	Name     string
	Price    float64
	Quantity int
}

type Bill struct {
	ID            uint
	PaymentID     uint
	Amount        float64
	Currency      string
	PaymentMethod string
	Products      []Product
	IssueDate     time.Time
}
