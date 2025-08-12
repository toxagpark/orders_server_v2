package repository

import "WB_LVL_0_NEW/internal/domain/model"

type OrderValidator interface {
	Validate(order *model.Order) error
}
