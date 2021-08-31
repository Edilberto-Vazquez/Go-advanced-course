package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) Getid() int {
	return e.id
}

func (e Employee) Getname() string {
	return e.name
}

func main() {
	e1 := Employee{id: 1, name: "Edilberto"}
	fmt.Printf("%v\n", e1)

	// methods
	e1.SetId(2)
	e1.SetName("Maria")
	fmt.Printf("%v\n", e1)
}
