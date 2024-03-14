package ports

import (
	"context"

	"ApuestaTotal/internal/payment/domain/dto"
	"ApuestaTotal/internal/payment/domain/entity"
)

type Payment interface {
	Create(ctx context.Context, newPayment dto.PaymentCreate) (entity.Payment, error)
	GetById(ctx context.Context, id int) (entity.Payment, error)
}
