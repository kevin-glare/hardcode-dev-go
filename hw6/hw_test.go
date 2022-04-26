package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{
			name: "#1",
			args: []float64{1, 1, 4, 5},
			want: 5,
		},
		{
			name: "#2",
			args: []float64{1, 1, 4, -5},
			want: -1,
		},
		{
			name: "#3",
			args: []float64{0, 0, 0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDistance(tt.args[0], tt.args[1], tt.args[2], tt.args[3]) // очень не нравится данный подход, жалко не работает tt.geom...
			if got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
