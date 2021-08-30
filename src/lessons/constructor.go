// package lessons

// import "fmt"

// type EmployeeC struct {
// 	id        int
// 	name      string
// 	vacations bool
// }

// func NewEmployee(id int, name string, vacations bool) *EmployeeC {
// 	return &EmployeeC{
// 		id:        id,
// 		name:      name,
// 		vacations: vacations,
// 	}
// }

// func Constructor() {
// 	// form 1
// 	e := EmployeeC{}
// 	fmt.Printf("%v\n", e)

// 	// form 2
// 	e2 := EmployeeC{
// 		id:        1,
// 		name:      "Edilberto",
// 		vacations: true,
// 	}
// 	fmt.Printf("%v\n", e2)

// 	// form 3
// 	e3 := new(EmployeeC)
// 	fmt.Printf("%v\n", *e3)
// 	e3.id = 15
// 	e3.name = "Maria"
// 	e3.vacations = false
// 	fmt.Printf("%v\n", *e3)

// 	// form4
// 	e4 := NewEmployee(1, "Name", false)
// 	fmt.Printf("%v\n", *e4)
// }
