package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	infra "WB_LVL_0_NEW/internal/infrastructure/validator"
)

func NewValidate() repository.OrderValidator {
	return infra.NewOrderValidator()
}
