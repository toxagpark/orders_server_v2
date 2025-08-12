package repository

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
)

type simpleConsumer struct {
	consumer sarama.Consumer
	topic    string
}

func NewSimpleConsumer(consumer sarama.Consumer, topic string) *simpleConsumer {
	return &simpleConsumer{
		consumer: consumer,
		topic:    topic,
	}
}

func (c *simpleConsumer) StartConsuming(ctx context.Context, handler func(ctx context.Context, order model.Order) error) error {
	pc, err := c.consumer.ConsumePartition(c.topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer pc.Close()

	for {
		select {
		case msg := <-pc.Messages():
			log.Println("New message")
			order := model.Order{}
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("JSON parse error: %v | Message: %s", err, string(msg.Value))
				continue
			}

			var validate *validator.Validate = validator.New()
			if err := validate.Struct(order); err != nil {
				log.Printf("Validation error: %v | Message: %s", err, string(msg.Value))
				continue
			}

			handler(ctx, order)
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *simpleConsumer) Close() error {
	return c.consumer.Close()
}
