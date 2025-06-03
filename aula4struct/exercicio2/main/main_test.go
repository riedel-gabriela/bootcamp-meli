package main

import "testing"

func TestPrintEmployee(t *testing.T) {
	tests := []struct {
		ID          int
		Name        string
		DateOfBirth string
		Position    string
		expected    string
	}{
		{1, "John Doe", "1990-01-01", "Software Engineer", "ID: 1, Name: John Doe, Position: Software Engineer, Date of Birth: 1990-01-01"},
	}

	for _, test := range tests {
		person := Person{
			ID:          test.ID,
			Name:        test.Name,
			DateOfBirth: test.DateOfBirth,
		}
		employee := Employee{
			ID:       test.ID,
			Position: test.Position,
			Person:   person,
		}
		result := employee.PrintEmployee()
		if result != test.expected {
			t.Errorf("PrintEmployee() = %s; want %s", result, test.expected)
		}
	}

}
