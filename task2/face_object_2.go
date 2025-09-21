package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println(e.EmployeeID, e.Person.Name, e.Person.Age)
}

func main() {
	e := Employee{
		Person: Person{
			Name: "keen",
			Age:  20,
		},
		EmployeeID: 1,
	}
	e.PrintInfo()
}
