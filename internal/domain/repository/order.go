package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
	GetByUID(ctx context.Context, uid string) (model.Order, error)
	Close() error
}
