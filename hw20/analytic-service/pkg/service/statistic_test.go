package service

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/link"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"testing"
)

func TestStatistic_Update(t *testing.T) {
	type args struct {
		link model.Link
	}
	tests := []struct {
		name string
		args model.Link
		want Data
	}{
		{
			name: "Test #1",
			args: model.Link{primitive.ObjectID{}, "http://go.com", link.ShortLink()},
			want: Data{1, 6, 6},
		},
		{
			name: "Test #2",
			args: model.Link{primitive.ObjectID{}, "https://vk.ru", link.ShortLink()},
			want: Data{2, 5.5, 11},
		},
	}

	s := &Statistic{
		Mutex: sync.Mutex{},
		Data:  Data{Count: 0, MidLength: 0, Length: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Update(tt.args); err != nil {
				t.Errorf("Update() error = %s", err.Error())
			}

			if s.Count != tt.want.Count {
				t.Errorf("Update() error = %v, want %v", s.Count, tt.want)
			}

			if s.MidLength != tt.want.MidLength {
				t.Errorf("Update() error = %v, want %v", s.MidLength, tt.want.MidLength)
			}

			if s.Length != tt.want.Length {
				t.Errorf("Update() error = %v, want %v", s.Length, tt.want.Length)
			}
		})
	}
}
