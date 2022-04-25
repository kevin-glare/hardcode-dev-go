package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         []float64
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         []float64{1, 1, 4, 5},
			wantDistance: 5,
		},
		{
			name:         "#2",
			geom:         []float64{1, 1, 4, -5},
			wantDistance: -1,
		},
		{
			name:         "#3",
			geom:         []float64{0, 0, 0, 0},
			wantDistance: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance := CalculateDistance(tt.geom[0], tt.geom[1], tt.geom[2], tt.geom[3]) // очень не нравится данный подход, жалко не работает tt.geom...
			if distance != tt.wantDistance {
				t.Errorf("CalculateDistance() = %v, want %v", distance, tt.wantDistance)
			}
		})
	}
}