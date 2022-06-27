package service

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"strings"
	"sync"
)

type Data struct {
	Count     int     `json:"count"`
	MidLength float64 `json:"mid_length"`
	Length    int     `json:"length"`
}

type Statistic struct {
	sync.Mutex
	Data
}

func NewStatistic() *Statistic {
	return &Statistic{}
}

func (s *Statistic) Update(link model.Link) error {
	s.Lock()
	defer s.Unlock()

	s.Count++
	s.Length += len(strings.Split(link.Url, "//")[1])
	s.MidLength = float64(s.Length) / float64(s.Count)

	return nil
}

func (s *Statistic) StatisticData() Data {
	s.Lock()
	defer s.Unlock()

	return s.Data
}
