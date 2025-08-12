package main

import (
	"WB_LVL_0_NEW/internal/domain/services"
	"WB_LVL_0_NEW/internal/infrastructure/config"
	router "WB_LVL_0_NEW/internal/interfaces/http"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting application...")

	// Инициализация зависимостей
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer db.Close()

	cache, err := config.NewRedis()
	if err != nil {
		log.Fatalf("Redis error: %v", err)
	}
	defer cache.Close()

	// Создание сервиса
	validator := config.NewValidate()
	service := services.NewOrderService(db, cache, validator)

	// Настройка Kafka consumer
	consumer, err := config.NewKafkaConfig().NewConsumer()
	if err != nil {
		log.Fatalf("Kafka error: %v", err)
	}

	// Контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запуск consumer в горутине
	go func() {
		if err := consumer.StartConsuming(ctx, service.HandleOrderCreated); err != nil {
			log.Printf("Consumer error: %v", err)
			cancel() // Инициируем shutdown при ошибке
		}
	}()

	// HTTP сервер
	router := router.NewRouter(service)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Router,
	}

	// Запуск сервера в горутине
	go func() {
		log.Println("Server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			cancel() // Инициируем shutdown при ошибке
		}
	}()

	// Ожидание сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gracefully...")

	// Остановка сервера с таймаутом
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// Отмена контекста остановит consumer
	cancel()
	time.Sleep(time.Second * 5)

	log.Println("Application stopped")
}
