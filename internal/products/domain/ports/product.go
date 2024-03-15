package ports

import (
	"context"

	"ApuestaTotal/internal/products/domain/dto"
	"ApuestaTotal/internal/products/domain/entity"
)

type Product interface {
	Create(ctx context.Context, newProduct dto.ProductCreate) (entity.Product, error)
	Update(ctx context.Context, updateProduct dto.ProductUpdate) error
	GetById(ctx context.Context, id uint) (entity.Product, error)
	GetAll(ctx context.Context) ([]entity.Product, error)
}
