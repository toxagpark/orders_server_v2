package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/repository"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
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
		return c.redisClient.Close()
	}
	return nil
}

func (c *cacheOrderRepo) Set(order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return c.redisClient.Set(order.OrderUID, data, c.ttl).Err()
}

func (c *cacheOrderRepo) Get(uid string) (model.Order, error) {
	data, err := c.redisClient.Get(uid).Bytes()
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return model.Order{}, err
	}

	return order, nil
}
