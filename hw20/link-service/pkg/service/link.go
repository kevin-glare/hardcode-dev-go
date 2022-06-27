package service

import (
	"context"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/model"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/repository"
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

func (s *LinkService) NewLink(ctx context.Context, url string) (string, error) {
	err := s.producer.SendMessage(ctx, url)
	if err == nil {
		return "", err
	}

	link, err := s.repo.FindLink(ctx, url)
	if err == nil {
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
