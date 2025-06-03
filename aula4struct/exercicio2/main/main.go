package main

import "fmt"

func main() {
	myPerson := Person{
		ID:          1,
		Name:        "John Doe",
		DateOfBirth: "1990-01-01",
	}
	myEmployee := Employee{
		ID:       1,
		Position: "Software Engineer",
		Person:   myPerson,
	}
	println(myEmployee.PrintEmployee())
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e *Employee) PrintEmployee() string {
	return "ID: " + fmt.Sprint(e.ID) + ", Name: " + e.Name + ", Position: " + e.Position + ", Date of Birth: " + e.DateOfBirth
}
