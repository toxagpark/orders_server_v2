package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/validation"
)

func NewValidate() repository.OrderValidator {
	return validation.NewOrderValidator()
}
