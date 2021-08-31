package main

import "fmt"

func sum(numbers ...int) (total int) {
	for _, v := range numbers {
		total += v
	}
	return
}

func unnamedReturn(names ...string) {
	for _, v := range names {
		fmt.Println("Hello", v)
	}
}

func getValues() (double float64, integer int, characters string) {
	double = 2.5 * 2.5
	integer = 2 * 2
	characters = "Unnamed Returns"
	return
}

func main() {
	println(sum(1, 2, 3, 4, 5, 6))
	unnamedReturn("Edilberto", "Maria", "Rosa", "Juan")
	println(getValues())
}
