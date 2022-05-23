package task2

import (
	"github.com/kevin-glare/hardcode-dev-go/hw9/pkg/task1"
)

func MaxAgeHuman(people ...task1.Interface) *task1.Interface {
	var human *task1.Interface
	var maxAge int

	for _, p := range people {
		if p.Age() > maxAge {
			maxAge = p.Age()
			human = &p
		}
	}

	return human
}