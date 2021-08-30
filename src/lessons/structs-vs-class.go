package lessons

import "fmt"

type Employee struct {
	id   int
	name string
}

func MakeObject() {
	e := Employee{}
	e2 := Employee{id: 2, name: "Maria"}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "Edilberto"
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", e2)
}
