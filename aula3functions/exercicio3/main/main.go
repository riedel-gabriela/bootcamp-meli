package main

/*
Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria.

Categoria C, o salário deles é de US$ 1.000 por hora.
Categoria B, o salário deles é de US$ 1.500 por hora, mais 20% do salário mensal.
Categoria A, o salário deles é de US$ 3.000 por hora, mais 50% do salário mensal.
Você deve gerar uma função que receba como parâmetro o número de minutos trabalhados por mês, a categoria e retorne o salário.
*/
import "fmt"

func main() {
	fmt.Println("Salario categoria A: U$", somaSalario(10560, "A"))
	fmt.Println("Salario categoria B: U$", somaSalario(10560, "B"))
	fmt.Println("Salario categoria C: U$", somaSalario(10560, "C"))
}

func somaSalario(minutosTrabalhados int, categoria string) (salario float64) {
	switch categoria {
	case "C":
		salario = (float64(minutosTrabalhados) / 60) * 1000.0
	case "B":
		salario = ((float64(minutosTrabalhados) / 60) * 1500.0) + 1.2
	case "A":
		salario = ((float64(minutosTrabalhados) / 60) * 3000.0) + 1.5
	}
	return
}
