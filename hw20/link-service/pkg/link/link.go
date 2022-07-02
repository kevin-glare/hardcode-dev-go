package link

import (
	"context"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	n       = 8
)

type Struct struct {
	repo     *repository.LinkRepo
	producer *kfk.Producer
}

func New(repo *repository.LinkRepo, producer *kfk.Producer) *Struct {
	return &Struct{repo: repo, producer: producer}
}

func (s *Struct) Link(ctx context.Context, short_url string) (*model.Link, error) {
	return s.repo.FindLink(ctx, bson.M{"short_url": short_url})
}

func (s *Struct) NewLink(ctx context.Context, url string) (string, error) {
	link, err := s.repo.FindLink(ctx, bson.M{"url": url})
	if err == nil {
		go s.producer.SendMessage(ctx, *link)
		return link.ShortUrl, nil
	}

	link = &model.Link{
		Url:      url,
		ShortUrl: ShortLink(),
	}

	err = s.repo.CreateLink(ctx, link)
	if err != nil {
		return "", err
	}

	go s.producer.SendMessage(ctx, *link)

	return link.ShortUrl, nil
}

func ShortLink() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
