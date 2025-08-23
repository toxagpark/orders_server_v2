package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"

	"context"
)

type Consumer interface {
	StartConsuming(ctx context.Context, handler func(ctx context.Context, order model.Order) error) error
	Close() error
}
