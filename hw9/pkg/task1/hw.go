package task1

type Interface interface {
	Age() int
}

type Customer struct {
	age int
}

type Employee struct {
	age int
}

func NewCustomer(age int) *Customer {
	return &Customer{age: age}
}

func NewEmployee(age int) *Employee {
	return &Employee{age: age}
}

func (c *Customer) Age() int {
	return c.age
}
func (e *Employee) Age() int {
	return e.age
}

func MaxAge(people ...Interface) int {
	var maxAge int
	for _, p := range people {
		if p.Age() > maxAge {
			maxAge = p.Age()
		}
	}

	return maxAge
}
