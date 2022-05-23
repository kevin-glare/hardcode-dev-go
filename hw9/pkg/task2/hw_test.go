package task2

import (
	"testing"

	"github.com/kevin-glare/hardcode-dev-go/hw9/pkg/task1"
)

func TestMaxAgeHuman(t *testing.T) {
	// Не нравится такой подход, пробовал testify.Assert и reflect.DeepEqual, но из-за указателей тесты не проходят (MaxAgeHuman() = &{2}, want &{2}).

	employee1 := task1.NewEmployee(1)
	employee2 := task1.NewEmployee(2)
	customer1 := task1.NewCustomer(1)
	customer2 := task1.NewCustomer(2)

	tests := []struct {
		name string
		args []task1.Interface
		want task1.Interface
	}{
		{
			name: "Test #1 - all types",
			args: []task1.Interface{
				customer1,
				employee2,
			},
			want: employee2,
		},
		{
			name: "Test #2 - only employees",
			args: []task1.Interface{
				customer1,
				customer2,
			},
			want: customer2,
		},
		{
			name: "Test #3 - only employees",
			args: []task1.Interface{
				employee1,
				employee2,
			},
			want: employee2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAgeHuman(tt.args...)

			if *got != tt.want {
				t.Errorf("MaxAgeHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
