package config

import (
	"WB_LVL_0_NEW/internal/domain/repository"
	kafka "WB_LVL_0_NEW/internal/infrastructure/kafka/repository"

	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
	}
}

func (kc *KafkaConfig) NewConsumer() (repository.Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(kc.Brokers, config)
	if err != nil {
		return nil, err
	}

	return kafka.NewSimpleConsumer(consumer, kc.Topic), nil
}
