package cache

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"

	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	ErrJSON       = errors.New("json error")
	ErrRedisSet   = errors.New("redis set error")
	ErrRedisGet   = errors.New("redis get error")
	ErrRedisClose = errors.New("redis close error")
)

type cacheOrderRepo struct {
	redisClient *redis.Client
	ttl         time.Duration
}

func NewCacheOrderRepository(client *redis.Client, ttl time.Duration) repository.CacheOrderRepository {
	return &cacheOrderRepo{redisClient: client, ttl: ttl}
}

func (c *cacheOrderRepo) Close() error {
	if c.redisClient != nil {
		err := c.redisClient.Close()
		if err != nil {
			return fmt.Errorf("%w: %w", ErrRedisClose, err)
		}
	}
	return nil
}

func (c *cacheOrderRepo) Set(order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrJSON, err)
	}
	err = c.redisClient.Set(order.OrderUID, data, c.ttl).Err()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRedisSet, err)
	}

	return nil
}

func (c *cacheOrderRepo) Get(uid string) (model.Order, error) {
	data, err := c.redisClient.Get(uid).Bytes()
	if err != nil {
		return model.Order{}, fmt.Errorf("%w: %w", ErrRedisGet, err)
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return model.Order{}, fmt.Errorf("%w: %w", ErrJSON, err)
	}

	return order, nil
}
