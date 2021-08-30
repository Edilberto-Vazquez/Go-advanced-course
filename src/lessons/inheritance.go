// package lessons

// import "fmt"

// type PrintInfo interface {
// 	getMessage() string
// }

// type PersonIn struct {
// 	name string
// 	age  int
// }

// type EmployeeIn struct {
// 	id int
// }

// type FullTimeEmployee struct {
// 	PersonIn
// 	EmployeeIn
// 	endDate string
// }

// func (ftEmployee FullTimeEmployee) getMessage() string {
// 	return "Full Time Employee"
// }

// type TemporaryEmployee struct {
// 	PersonIn
// 	Employee
// 	textRate int
// }

// func getMessage(p PrintInfo) {
// 	fmt.Println(p.getMessage())
// }

// func Inheritance() {
// 	// ftEmployee := FullTimeEmployee{}
// 	// ftEmployee.name = "Name"
// 	// ftEmployee.age = 2
// 	// ftEmployee.id = 16549
// 	// fmt.Printf("%v\n", ftEmployee)
// 	// tEmployee := TemporaryEmployee{}
// 	// getMessage(tEmployee)
// 	// getMessage(ftEmployee)
// }
