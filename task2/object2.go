package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	employeeId int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("name: %s, age: %d, EmployeeId: %d\n", e.name, e.age, e.employeeId)
}

func main() {
	employee := Employee{
		Person{"poroon", 28}, 9527,
	}
	employee.PrintInfo()
}
