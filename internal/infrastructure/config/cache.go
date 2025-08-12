package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	redisRepo "WB_LVL_0_NEW/internal/infrastructure/redis/repository"
	"time"

	"github.com/go-redis/redis"
)

func NewRedis() (repository.CacheOrderRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ttl := 10 * time.Minute

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisRepo.NewCacheOrderRepository(client, ttl), nil
}
