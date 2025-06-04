package main

import (
	"errors"
	"fmt"
)

/*
Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int". Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: salary is less than 10000" e a inicie caso "salary" seja menor ou igual a 10000. A validação deve ser feita com a função "Is()" dentro do "main".
*/

// custom error type para salários menores ou iguais a 10000
type LowSalaryError struct{}

// Error implementa a interdace de erro para LowSalaryError.
func (e *LowSalaryError) Error() string {
	return "Error: salary is less than 10000"
}

func main() {
	salary := 5000

	if salary <= 10000 {
		// cria uma instancia do erro personalizado
		errReturned := &LowSalaryError{}
		if errors.Is(errReturned, &LowSalaryError{}) {
			fmt.Println(errReturned)
		}
	} else {
		fmt.Println("Salary is valid")
	}
}
