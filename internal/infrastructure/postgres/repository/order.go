package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/postgres/converter"
	"WB_LVL_0_NEW/internal/infrastructure/postgres/dto"
	"context"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(ctx context.Context, order model.Order) error {
	dbOrder := converter.ToDTO(order)
	return r.db.WithContext(ctx).Create(&dbOrder).Error
}

func (r *orderRepo) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (r *orderRepo) GetByUID(ctx context.Context, uid string) (model.Order, error) {
	var dbOrder dto.Order
	err := r.db.WithContext(ctx).
		Preload("Items").
		First(&dbOrder, "order_uid = ?", uid).Error

	if err != nil {
		return model.Order{}, err
	}

	return converter.ToDomain(dbOrder), nil
}
