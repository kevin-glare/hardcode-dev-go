package service

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

type LinkService struct {
	repo     *repository.LinkRepo
	producer *kfk.Producer
}

func NewLinkService(repo *repository.LinkRepo, producer *kfk.Producer) *LinkService {
	return &LinkService{repo: repo, producer: producer}
}

func (s *LinkService) Link(ctx context.Context, short_url string) (*model.Link, error) {
	return s.repo.FindLink(ctx, bson.M{"short_url": short_url})
}

func (s *LinkService) NewLink(ctx context.Context, url string) (string, error) {
	link, err := s.repo.FindLink(ctx, bson.M{"url": url})
	if err == nil {
		go s.producer.SendMessage(ctx, *link)
		return link.ShortUrl, nil
	}

	link = &model.Link{
		Url:      url,
		ShortUrl: shortLink(),
	}

	err = s.repo.CreateLink(ctx, link)
	if err != nil {
		return "", err
	}

	go s.producer.SendMessage(ctx, *link)

	return link.ShortUrl, nil
}

func shortLink() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
