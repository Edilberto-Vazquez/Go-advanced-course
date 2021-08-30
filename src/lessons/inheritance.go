package lessons

import "fmt"

type PersonIn struct {
	name string
	age  int
}

type EmployeeIn struct {
	id int
}

type FullTimeEmployee struct {
	PersonIn
	EmployeeIn
}

func GetMessage(p PersonIn) {
	fmt.Printf("%s With age %d\n", p.name, p.age)
}

func Inheritance() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 16549

	fmt.Printf("%v", ftEmployee)
}
