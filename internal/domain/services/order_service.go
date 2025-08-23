package services

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/cache"
	"WB_LVL_0_NEW/internal/infrastructure/order"
	"WB_LVL_0_NEW/internal/infrastructure/validation"

	"context"
	"errors"
	"fmt"
	"log"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type OrderService struct {
	orderRepo    repository.OrderRepository
	cachRepo     repository.CacheOrderRepository
	validateRepo repository.OrderValidator
}

func NewOrderService(db repository.OrderRepository, cache repository.CacheOrderRepository, validate repository.OrderValidator) *OrderService {
	return &OrderService{
		orderRepo:    db,
		cachRepo:     cache,
		validateRepo: validate,
	}
}

func (s *OrderService) HandleOrderCreated(ctx context.Context, orderModel model.Order) error {
	if err := s.validateRepo.Validate(&orderModel); errors.Is(err, validation.ErrValidate) {
		return err
	}
	if err := s.orderRepo.Create(ctx, orderModel); errors.Is(err, order.ErrCreateDB) {
		return err
	}
	if err := s.cachRepo.Set(orderModel); errors.Is(err, cache.ErrRedisSet) ||
		errors.Is(err, cache.ErrJSON) {
		log.Printf("cache error: %v\n", err)
	}
	return nil
}

func (s *OrderService) HandleOrderGet(ctx context.Context, orderUID string) (*model.Order, error) {
	orderModel, err := s.cachRepo.Get(orderUID)
	if errors.Is(err, cache.ErrRedisGet) || errors.Is(err, cache.ErrJSON) {
		if errors.Is(err, cache.ErrJSON) {
			log.Printf("cache error: %v\n", err)
		}
		orderModel, err = s.orderRepo.GetByUID(ctx, orderUID)
		if errors.Is(err, order.ErrGetDB) {
			return nil, fmt.Errorf("%w:, %w", err, ErrOrderNotFound)
		}
	}
	err = s.cachRepo.Set(orderModel)
	if errors.Is(err, cache.ErrRedisSet) || errors.Is(err, cache.ErrJSON) {
		log.Printf("cache error: %v", err)
	}
	return &orderModel, nil
}
