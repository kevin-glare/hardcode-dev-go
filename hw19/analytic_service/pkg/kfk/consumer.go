package kfk

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type Statistic struct {
	Count     int
	MidLength float64
	Length    int
}

type Analytic struct {
	Reader *kafka.Reader

	sync.Mutex
	Statistic
}

func NewConsumer(brokers []string, topic, groupID string) (*Analytic, error) {
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

	return &Analytic{
		Reader:    r,
		Statistic: Statistic{},
	}, nil
}

func (s *Analytic) ConsumerRun() {
	log.Println("ConsumerRun")
	for {
		msg, err := s.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println(string(msg.Value))

		s.Lock()
		s.Count++
		s.Length += len(string(msg.Value))
		s.MidLength = float64(s.Length) / float64(s.Count)
		s.Unlock()

		err = s.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}
