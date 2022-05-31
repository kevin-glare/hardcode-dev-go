package task2

type Customer struct {
	age int
}

type Employee struct {
	age int
}

func NewCustomer(age int) Customer {
	return Customer{age: age}
}

func NewEmployee(age int) Employee {
	return Employee{age: age}
}

func MaxAgeHuman(people []interface{}) interface{} {
	var human interface{}
	var maxAge int

	for _, el := range people {
		if p, ok := el.(Customer); ok {
			if p.age > maxAge {
				maxAge = p.age
				human = p
			}
		}

		if p, ok := el.(Employee); ok {
			if p.age > maxAge {
				maxAge = p.age
				human = p
			}
		}
	}

	return human
}
