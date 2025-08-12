package services

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"
	"context"
	"errors"
	"log"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type OrderService struct {
	dbRepo       repository.OrderRepository
	cachRepo     repository.CacheOrderRepository
	validateRepo repository.OrderValidator
}

func NewOrderService(db repository.OrderRepository, cache repository.CacheOrderRepository, validate repository.OrderValidator) *OrderService {
	return &OrderService{
		dbRepo:       db,
		cachRepo:     cache,
		validateRepo: validate,
	}
}

func (s *OrderService) HandleOrderCreated(ctx context.Context, order model.Order) error {
	if err := s.validateRepo.Validate(&order); err != nil {
		return err
	}
	if err := s.dbRepo.Create(ctx, order); err != nil {
		return err
	}
	if err := s.cachRepo.Set(order); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) HandleOrderGet(ctx context.Context, orderUID string) (*model.Order, error) {
	order, err := s.cachRepo.Get(orderUID)
	if err != nil {
		order, err = s.dbRepo.GetByUID(ctx, orderUID)
		if err != nil {
			return nil, err
		}
	}
	err = s.cachRepo.Set(order)
	if err != nil {
		log.Printf("failed to cache order (UID: %s): %v", orderUID, err)
	}
	return &order, nil
}
