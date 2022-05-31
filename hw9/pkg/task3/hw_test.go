package task3

import (
	"strings"
	"testing"
)

func TestPrintString(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{
			name: "Test #1 - one string",
			args: []interface{}{"A"},
			want: "A",
		},
		{
			name: "Test #2 - strings",
			args: []interface{}{"A", " ", "B"},
			want: "A B",
		},
		{
			name: "Test #3 - strings and int",
			args: []interface{}{"A", " ", "B", 1},
			want: "A B",
		},
		{
			name: "Test #4 - strings and nil",
			args: []interface{}{"A", " ", "B", nil},
			want: "A B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := new(strings.Builder)

			PrintString(w, tt.args...)

			if got := w.String(); got != tt.want {
				t.Errorf("Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
