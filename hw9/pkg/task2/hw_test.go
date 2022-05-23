package task2

import (
	"testing"

	"github.com/kevin-glare/hardcode-dev-go/hw9/pkg/task1"
	"gotest.tools/assert"
)

func TestMaxAgeHuman(t *testing.T) {
	tests := []struct {
		name string
		args []task1.Interface
		want task1.Interface
	}{
		{
			name: "Test #1 - customer and employee",
			args: []task1.Interface{
				task1.NewCustomer(1),
				task1.NewEmployee(2),
			},
			want: task1.NewEmployee(2),
		},
		{
			name: "Test #2 - only customers",
			args: []task1.Interface{
				task1.NewCustomer(1),
				task1.NewCustomer(2),
			},
			want: task1.NewCustomer(2),
		},
		{
			name: "Test #3 - only employees",
			args: []task1.Interface{
				task1.NewEmployee(1),
				task1.NewEmployee(2),
			},
			want: task1.NewEmployee(2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAgeHuman(tt.args...)

			assert.Equal(t, tt.want, got)
		})
	}
}
