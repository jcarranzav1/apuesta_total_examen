package model

import (
	"ApuestaTotal/internal/cart/domain/entity"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Items []Item `gorm:"constraint:	onDelete:CASCADE"`
}

type Item struct {
	gorm.Model
	ProductID uint
	Quantity  int
	CartID    uint
}

func (p *Cart) ToProductDomain() entity.Cart {
	var domainProducts []entity.Item
	for _, item := range p.Items {
		domainProducts = append(domainProducts, entity.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return entity.Cart{
		ID:    p.ID,
		Items: domainProducts,
	}
}
