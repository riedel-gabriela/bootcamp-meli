package main

/*
Uma universidade precisa registrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada aluno,
como segue:

Nome: [Primeiro nome do aluno]
Sobrenome: [Sobrenome do aluno]
ID: [ID do aluno]
Data: [Data de admissão do aluno

Os valores entre colchetes devem ser substituídos pelos dados fornecidos pelos alunos.
Para isso, é necessário gerar uma estrutura Students com as variáveis Name, Surname, DNI, Date
e que tenha um método de detalhamento
*/

import "fmt"

type Student struct {
	DNI     int
	Name    string
	Surname string
	Date    string
}

type StudentInterface interface {
	NewStudent(name, surname string, dni int, date string) Student
	PrintDetails() string
}

func main() {

	student := Student{
		Name:    "João",
		Surname: "Silva",
		DNI:     123456,
		Date:    "2023-01-01",
	}

	fmt.Println(student.PrintDetails())

	newStudent := student.NewStudent("Maria", "Oliveira", 654321, "2023-02-01")
	fmt.Println(newStudent.PrintDetails())

}

func (s Student) PrintDetails() string {
	return fmt.Sprintf("Nome: %s, Sobrenome: %s, ID: %d, Data: %s", s.Name, s.Surname, s.DNI, s.Date)
}

func (s Student) NewStudent(name, surname string, dni int, date string) Student {
	return Student{
		Name:    name,
		Surname: surname,
		DNI:     dni,
		Date:    date,
	}
}
