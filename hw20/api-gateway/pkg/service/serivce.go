package service

import (
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/delivery"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"net/url"
	"os"
)

type ApiService struct {
	producer *kfk.Producer
}

func NewApiService(producer *kfk.Producer) *ApiService {
	return &ApiService{
		producer: producer,
	}
}

func (s *ApiService) Statistic() (delivery.ParsingResponse, error) {
	url := fmt.Sprintf("http://localhost%s/api/v1/statistic", os.Getenv("ANALYTIC_SERVICE_HTTP_HOST"))
	return delivery.HttpGet(url)
}

func (s *ApiService) ShowLink(shortURL string) (delivery.ParsingResponse, error) {
	url := fmt.Sprintf("http://localhost%s/api/v1/links/%s", os.Getenv("CACHE_SERVICE_HTTP_HOST"), shortURL)
	res, err := delivery.HttpGet(url)
	if err == nil {
		return res, nil
	}

	url = fmt.Sprintf("http://localhost%s/api/v1/links/%s", os.Getenv("LINK_SERVICE_HTTP_HOST"), shortURL)
	return delivery.HttpGet(url)
}

func (s *ApiService) AddLink(linkURL string) (delivery.ParsingResponse, error) {
	uri := fmt.Sprintf("http://localhost%s/api/v1/links", os.Getenv("LINK_SERVICE_HTTP_HOST"))

	payload := url.Values{
		"url": {linkURL},
	}

	return delivery.HttpPost(uri, payload)
}
