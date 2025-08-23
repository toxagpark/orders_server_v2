package events

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator"
)

var (
	ErrConsuming = errors.New("error consuming")
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
		return fmt.Errorf("%w: %w", ErrConsuming, err)
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

			if err = handler(ctx, order); err != nil {
				log.Printf("handler err: %v", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *simpleConsumer) Close() error {
	return c.consumer.Close()
}
