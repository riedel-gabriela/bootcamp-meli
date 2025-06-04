package main

import (
	"errors"
	"fmt"
)

/*
Faça o mesmo que no exercício anterior, mas reformule o código de modo que, em vez de "Error()", seja implementado "errors.New()".
*/
var err = errors.New("error: salary is less than 10000")

func main() {
	salary := 5000

	err_returned := verifySalary(salary)

	if errors.Is(err_returned, err) {
		fmt.Println(err)
	} else {
		fmt.Println("Salary is valid")
	}
}

func verifySalary(salary int) error {
	if salary <= 10000 {
		return err
	}
	return nil
}
