package service

import "github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"

type ApiService struct {
	producer *kfk.Producer
}

func NewApiService(producer *kfk.Producer) *ApiService {
	return &ApiService{
		producer: producer,
	}
}

func (s *ApiService) Statistic() {

}
