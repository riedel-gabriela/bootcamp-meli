package main

import (
	"testing"
)

func TestPrintDetails(t *testing.T) {
	student := Student{
		Name:    "João",
		Surname: "Silva",
		DNI:     123456,
		Date:    "2023-01-01",
	}

	expected := "Nome: João, Sobrenome: Silva, ID: 123456, Data: 2023-01-01"
	if student.PrintDetails() != expected {
		t.Errorf("Expected %s but got %s", expected, student.PrintDetails())
	}

}
