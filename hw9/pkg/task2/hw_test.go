package task2

import (
	"testing"

	"github.com/kevin-glare/hardcode-dev-go/hw9/pkg/task1"
)

func TestMaxAgeHuman(t *testing.T) {
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
			name: "Test #1 - customer and employee",
			args: []task1.Interface{
				customer1,
				employee2,
			},
			want: employee2,
		},
		{
			name: "Test #2 - only customers",
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

			if tt.want != got {
				t.Errorf("MaxAgeHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
