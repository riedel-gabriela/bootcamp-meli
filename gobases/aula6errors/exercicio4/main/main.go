package main

import (
	"fmt"
)

/*
Repete o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de erro receba como parâmetro o valor de "salary" indicando que ele não atinge o mínimo tributável (a mensagem exibida pelo console deve dizer:“Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”, sendo  [salary] o valor do tipo int passado pelo parâmetro).
*/

func main() {
	salary := 10000

	if salary < 150000 {
		err := fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}
