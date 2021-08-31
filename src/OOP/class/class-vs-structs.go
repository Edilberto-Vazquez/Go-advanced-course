package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e1 := Employee{id: 1, name: "Edilberto"}
	e2 := Employee{}
	e2.id = 2
	e2.name = "Maria"
	fmt.Printf("%v\n", e1)
	fmt.Printf("%v\n", e2)
}
