package kfk

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	reader     *kafka.Reader
	handleFunc func(link model.Link) error
}

func NewConsumer(brokers []string, topic, groupID string, f func(link model.Link) error) (*Consumer, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" || groupID == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	return &Consumer{reader: r, handleFunc: f}, nil
}

func (c *Consumer) Run() {
	log.Println("Consumer Run")

	for {
		msg, err := c.reader.FetchMessage(context.Background())
		if err != nil {
			log.Printf("FetchMessage: %s", err.Error())
			continue
		}

		var link model.Link
		err = json.Unmarshal(msg.Value, &link)
		if err != nil {
			log.Printf("Unmarshal: %s", err.Error())
			continue
		}

		err = c.handleFunc(link)
		if err != nil {
			log.Println(err)
		}

		err = c.reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}
