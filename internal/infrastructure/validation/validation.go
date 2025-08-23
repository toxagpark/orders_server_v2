package validation

import (
	"WB_LVL_0_NEW/internal/domain/model"

	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidate = errors.New("error validate")
)

type OrderValidator struct {
	validate *validator.Validate
}

func NewOrderValidator() *OrderValidator {
	v := validator.New()
	return &OrderValidator{validate: v}
}

func (v *OrderValidator) Validate(order *model.Order) error {
	err := v.validate.Struct(order)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValidate, err)
	}

	return nil
}
