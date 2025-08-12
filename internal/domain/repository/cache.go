package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"
)

type CacheOrderRepository interface {
	Set(order model.Order) error
	Get(uid string) (model.Order, error)
	Close() error
}
