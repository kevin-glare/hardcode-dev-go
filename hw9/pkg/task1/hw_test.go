package task1

import "testing"

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name string
		args []Interface
		want int
	}{
		{
			name: "Test #1 - all types",
			args: []Interface{
				NewCustomer(1),
				NewEmployee(2),
			},
			want: 2,
		},
		{
			name: "Test #2 - only customers",
			args: []Interface{
				NewCustomer(1),
				NewCustomer(2),
			},
			want: 2,
		},
		{
			name: "Test #3 - only employees",
			args: []Interface{
				NewEmployee(1),
				NewEmployee(2),
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAge(tt.args...)
			if got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
