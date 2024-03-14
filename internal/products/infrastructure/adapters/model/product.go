package model

import (
	"ApuestaTotal/internal/products/domain/entity"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name  string
	Price float64
	Stock int
}

type MultipleProduct []Product

func (p *Product) ToProductDomain() entity.Product {
	return entity.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}

func (p MultipleProduct) ToProductDomainSlice() []entity.Product {
	var domainProducts []entity.Product

	for _, product := range p {
		domainProducts = append(domainProducts, product.ToProductDomain())
	}

	return domainProducts
}
