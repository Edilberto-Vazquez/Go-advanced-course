// package lessons

// import "fmt"

// type Employee struct {
// 	id   int
// 	name string
// }

// func (e *Employee) Setid(id int) {
// 	e.id = id
// }

// func (e *Employee) SetName(name string) {
// 	e.name = name
// }

// func (e *Employee) Getid() int {
// 	return e.id
// }

// func (e *Employee) GetName() string {
// 	return e.name
// }

// func MakeObject() {
// 	e := Employee{}
// 	fmt.Printf("%v\n", e)
// 	e.id = 1
// 	e.name = "Edilberto"
// 	fmt.Printf("%v\n", e)
// 	e.Setid(5)
// 	e.SetName("Maria")
// 	fmt.Printf("%v\n", e)
// 	fmt.Println(e.Getid())
// 	fmt.Println(e.GetName())
// }
