package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{Name: "Edilberto", Age: 31, DNI: "1"}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  32,
					DNI:  "1",
					Name: "Edilberto",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, v := range table {
		v.mockFunc()
		ft, err := GetFullTimeEmployeeById(v.id, v.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != v.expectedEmployee.Age {
			t.Errorf("Error, got %d expexted %d", ft.Age, v.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
