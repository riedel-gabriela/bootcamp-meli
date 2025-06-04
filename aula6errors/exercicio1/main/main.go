package main

import (
	"fmt"
)

/*
Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".

Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: the salary entered does not reach the taxable minimum" e acione-a caso "salary" seja menor que 150.000. Caso contrário, será necessário imprimir no console a mensagem "Must pay tax".
*/

func main() {
	salary := 10000

	err := fmt.Errorf("error: the salary entered does not reach the taxable minimum")

	if salary < 150000 {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}
