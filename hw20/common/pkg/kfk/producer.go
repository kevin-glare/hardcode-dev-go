package kfk

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	*kafka.Writer
}

func NewProducer(broker, topic string) (*Producer, error) {
	if len(broker) == 0 || topic == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{w}, nil
}

func (p *Producer) SendMessage(ctx context.Context, message string) error {
	msg := kafka.Message{
		Value: []byte(message),
	}

	return p.WriteMessages(ctx, msg)
}
