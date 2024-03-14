package ports

import (
	"context"

	"ApuestaTotal/internal/bill/domain/dto"
	"ApuestaTotal/internal/bill/domain/entity"
)

type Bill interface {
	Create(ctx context.Context, newBill dto.BillCreate) (entity.Bill, error)
	GetById(ctx context.Context, id int) (entity.Bill, error)
}
