package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	"WB_LVL_0_NEW/internal/infrastructure/events"

	"errors"
	"fmt"
	"os"

	"github.com/IBM/sarama"
)

var (
	ErrKafkaConsumer = errors.New("error kafka consumer")
)

type EventsConfig struct {
	Brokers []string
	Topic   string
}

func NewEventsConfig() *EventsConfig {
	return &EventsConfig{
		Brokers: []string{os.Getenv("EVENTS_ADDRESS")},
		Topic:   os.Getenv("EVENTS_TOPIC"),
	}
}

func (kc *EventsConfig) NewKafkaConsumer() (repository.Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(kc.Brokers, config)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrKafkaConsumer, err)
	}

	return events.NewSimpleConsumer(consumer, kc.Topic), nil
}
