package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	postgresDB "WB_LVL_0_NEW/internal/infrastructure/postgres/repository"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (repository.OrderRepository, error) {
	dsn, err := os.ReadFile("DBLog.txt")
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return postgresDB.NewOrderRepository(db), nil
}
