package order

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/order/converter"
	"WB_LVL_0_NEW/internal/infrastructure/order/dto"

	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrCreateDB = errors.New("error db insert order")
	ErrCloseDB  = errors.New("error close db")
	ErrGetDB    = errors.New("error get from db")
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(ctx context.Context, order model.Order) error {
	dbOrder := converter.ToDTO(order)
	if err := r.db.WithContext(ctx).Create(&dbOrder).Error; err != nil {
		return fmt.Errorf("%w: %w", ErrCreateDB, err)
	}
	return nil
}

func (r *orderRepo) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCloseDB, err)
	}
	if err = sqlDB.Close(); err != nil {
		return fmt.Errorf("%w: %w", ErrCloseDB, err)
	}
	return nil
}

func (r *orderRepo) GetByUID(ctx context.Context, uid string) (model.Order, error) {
	var dbOrder dto.Order
	err := r.db.WithContext(ctx).
		Preload("Items").
		First(&dbOrder, "order_uid = ?", uid).Error

	if err != nil {
		return model.Order{}, fmt.Errorf("%w: %w", ErrGetDB, err)
	}

	return converter.ToDomain(dbOrder), nil
}
