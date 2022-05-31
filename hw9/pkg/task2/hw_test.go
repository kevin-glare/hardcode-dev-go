package task2

import (
	"testing"
)

func TestMaxAgeHuman(t *testing.T) {
	employee1 := NewEmployee(1)
	employee2 := NewEmployee(2)
	customer1 := NewCustomer(1)
	customer2 := NewCustomer(2)

	tests := []struct {
		name string
		args []interface{}
		want interface{}
	}{
		{
			name: "Test #1 - customer and employee",
			args: []interface{}{
				customer1,
				employee2,
			},
			want: employee2,
		},
		{
			name: "Test #2 - only customers",
			args: []interface{}{
				customer1,
				customer2,
			},
			want: customer2,
		},
		{
			name: "Test #3 - only employees",
			args: []interface{}{
				employee1,
				employee2,
			},
			want: employee2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAgeHuman(tt.args)

			if tt.want != got {
				t.Errorf("MaxAgeHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
