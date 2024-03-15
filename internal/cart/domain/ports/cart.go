package ports

import (
	"context"

	"ApuestaTotal/internal/cart/domain/dto"
	"ApuestaTotal/internal/cart/domain/entity"
)

type Cart interface {
	CreateCart(ctx context.Context, newCart dto.CreateProductDTO) (entity.Cart, error)
	RemoveCart(ctx context.Context, id uint) (entity.Cart, error)
	AddProduct(ctx context.Context, addProduct dto.AddOrUpdateProductDTO) error
	RemoveProduct(ctx context.Context, removeProduct dto.RemoveProductDTO) error
	UpdateProduct(ctx context.Context, updateProduct dto.AddOrUpdateProductDTO) error
	GetById(ctx context.Context, id uint) (entity.Cart, error)
}
