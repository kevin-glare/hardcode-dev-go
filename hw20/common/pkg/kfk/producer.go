package kfk

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
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

func (p *Producer) SendMessage(ctx context.Context, message model.Link) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: payload,
	}

	return p.WriteMessages(ctx, msg)
}
