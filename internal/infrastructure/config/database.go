package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/order"

	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrPostgresClient = errors.New("error postgres client")
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func NewPostgresClient(cfg *DBConfig) (repository.OrderRepository, error) {
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPostgresClient, err)
	}
	return order.NewOrderRepository(db), nil
}

func buildDSN(cfg *DBConfig) string {
	return "host=" + cfg.Host +
		" user=" + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.Name +
		" port=" + cfg.Port +
		" sslmode=" + cfg.SSLMode
}
