package main

import "fmt"

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

type TemporaryEmployee struct {
	Person
	Employee
	textRate int
}

func main() {
	e := FullTimeEmployee{}
	e.name = "Edilberto"
	e.age = 23
	e.id = 6546
	e.endDate = "25/05/2021"
	fmt.Printf("%v\n", e)
}
