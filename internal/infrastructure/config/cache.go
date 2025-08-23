package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/cache"

	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	ErrCacheCfg    = errors.New("error cache cfg")
	ErrRedisClient = errors.New("error redis client")
)

type CacheConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewCacheConfig() (*CacheConfig, error) {
	dbStr := os.Getenv("CACHE_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCacheCfg, err)
	}
	return &CacheConfig{
		Addr:     os.Getenv("CACHE_ADDRESS"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       db,
	}, nil
}

func NewRedis(cfg *CacheConfig) (repository.CacheOrderRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ttl := 10 * time.Minute

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRedisClient, err)
	}
	return cache.NewCacheOrderRepository(client, ttl), nil
}
