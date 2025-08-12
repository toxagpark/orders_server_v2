package validator

import (
	"WB_LVL_0_NEW/internal/domain/model"

	"github.com/go-playground/validator/v10"
)

type OrderValidator struct {
	validate *validator.Validate
}

func NewOrderValidator() *OrderValidator {
	v := validator.New()
	return &OrderValidator{validate: v}
}

func (v *OrderValidator) Validate(order *model.Order) error {
	return v.validate.Struct(order)
}
