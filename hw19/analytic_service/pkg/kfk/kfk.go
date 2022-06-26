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
	Writer *kafka.Writer

	sync.Mutex
	Statistic
}

func New(brokers []string, topic, groupID string) (*Analytic, error) {
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

	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers[0]),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Analytic{
		Reader:    r,
		Writer:    w,
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

func (s *Analytic) SendMessage(message string) error {
	msg := kafka.Message{
		Value: []byte(message),
	}

	return s.Writer.WriteMessages(context.Background(), msg)
}
