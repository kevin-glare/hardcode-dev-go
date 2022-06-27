package service

import (
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
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

func (s *ApiService) Statistic() (api.ParsingResponse, error) {
	url := fmt.Sprintf("http://localhost%s/api/v1/statistic", os.Getenv("ANALYTIC_SERVICE_HTTP_HOST"))
	return api.Get(url)
}

func (s *ApiService) ShowLink(shortURL string) (api.ParsingResponse, error) {
	url := fmt.Sprintf("http://localhost%s/api/v1/links/%s", os.Getenv("CACHE_SERVICE_HTTP_HOST"), shortURL)
	res, err := api.Get(url)
	if err == nil {
		return res, nil
	}

	url = fmt.Sprintf("http://localhost%s/api/v1/links/%s", os.Getenv("LINK_SERVICE_HTTP_HOST"), shortURL)
	return api.Get(url)
}

func (s *ApiService) AddLink(linkURL string) (api.ParsingResponse, error) {
	url := fmt.Sprintf("http://localhost%s/api/v1/links", os.Getenv("LINK_SERVICE_HTTP_HOST"))
	payload := make(map[string]interface{})
	payload["url"] = linkURL

	return api.Post(url, payload)
}
