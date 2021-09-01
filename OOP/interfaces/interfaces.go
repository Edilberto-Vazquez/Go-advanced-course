package main

import "fmt"

type PrintInfo interface {
	getMessage() string
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ft FullTimeEmployee) getMessage() string {
	// fmt.Println("Fulltime employee")
	return "Fulltime Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	textRate int
}

func (te TemporaryEmployee) getMessage() string {
	// fmt.Println("Temporary Employee")
	return "Temporary Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	e := FullTimeEmployee{}
	e.name = "Edilberto"
	e.age = 23
	e.id = 6546
	e.endDate = "25/05/2021"
	fmt.Printf("%v\n", e)
	getMessage(e)
	e2 := TemporaryEmployee{}
	e2.name = "Edilberto"
	e2.age = 23
	e2.id = 6546
	e2.textRate = 10
	fmt.Printf("%v\n", e2)
	getMessage(e2)
}
