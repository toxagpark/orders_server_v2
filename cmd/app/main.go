package main

import (
	"WB_LVL_0_NEW/internal/domain/services"
	"WB_LVL_0_NEW/internal/handlers"
	"WB_LVL_0_NEW/internal/infrastructure/config"
	"WB_LVL_0_NEW/internal/infrastructure/events"

	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("Starting application...")

	// зависимости
	db_cfg := config.LoadDBConfig()
	db, err := config.NewPostgresClient(db_cfg)
	if errors.Is(err, config.ErrPostgresClient) {
		return err
	}
	defer db.Close()

	cacheCfg, err := config.NewCacheConfig()
	if errors.Is(err, config.ErrCacheCfg) {
		return err
	}
	cache, err := config.NewRedis(cacheCfg)
	if errors.Is(err, config.ErrRedisClient) {
		return err
	}
	defer cache.Close()

	validator := config.NewValidate()

	// Создание сервиса
	service := services.NewOrderService(db, cache, validator)

	// Настройка Kafka consumer
	consumer, err := config.NewEventsConfig().NewKafkaConsumer()
	if errors.Is(err, config.ErrKafkaConsumer) {
		return err
	}

	// HTTP сервер
	router := handlers.NewRouter(service)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router.Router,
	}

	// Контекст для graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// Запуск consumer в горутине
	go func() {
		if err := consumer.StartConsuming(ctx, service.HandleOrderCreated); errors.Is(err, events.ErrConsuming) {
			log.Println(err)
			stop()
		}
	}()

	// Запуск сервера в горутине
	go func() {
		log.Println("Server listening on :8081")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			stop()
		}
	}()

	// Ожидание сигнала завершения
	<-ctx.Done()

	log.Println("Shutting down gracefully...")

	// Остановка сервера с таймаутом
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Print(err)
	}

	log.Println("Application stopped")
	return nil
}
