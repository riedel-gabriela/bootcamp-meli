package main

import (
	"fmt"
)

/*
Vamos tornar nosso programa um pouco mais complexo e útil.

Desenvolver as funções necessárias para permitir que a empresa faça cálculos:
Salário mensal de um trabalhador de acordo com o número de horas trabalhadas. A função deve receber as horas trabalhadas no mês e o valor da hora como argumento.
A função deve retornar mais de um valor (salário calculado e erro). Caso o salário mensal seja igual ou superior a US$ 150.000, será deduzido um imposto de 10%. Caso o número de horas mensais inserido seja inferior a 80 ou um número negativo, a função deve retornar um erro. Ela indicará “Error: the worker cannot have worked less than 80 hours per month”.
*/

func main() {
	horasTrabalhadas := 85
	valorHora := 200.0

	salario, err := calculaSalario(horasTrabalhadas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("O salário calculado é: R$ %.2f\n", salario)

}

func calculaSalario(horasTrabalhadas int, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 {
		return 0, fmt.Errorf("error: the worker cannot have worked less than 80 hours per month")
	}

	salario := float64(horasTrabalhadas) * valorHora

	if salario >= 150000 {
		salario -= salario * 0.10
	}

	return salario, nil
}
