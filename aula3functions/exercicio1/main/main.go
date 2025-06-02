package main

/*
Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário.
Para cumprir o objetivo, é necessário criar uma função que retorne o imposto de um salário.
Levando em conta que se a pessoa ganha mais de US$ 50.000, 17% do salário será deduzido e se a pessoa ganha mais de US$ 150.000, 10% também será deduzido (27% no total).
*/
import "fmt"

func main() {
	// fmt.Println("Imposto sobre Salário: ", impostoSobreSalario(165000.0))
	fmt.Println("Imposto sobre Salário: ", impostoSobreSalario(51000.0))
	fmt.Println("Imposto sobre Salário: ", impostoSobreSalario(49000.0))
	fmt.Println("Imposto sobre Salário: ", impostoSobreSalario(151000.0))
}

func impostoSobreSalario(salario float64) float64 {
	result := 0.0
	switch {
	case salario <= 50000.0:
		break
	case salario <= 150000.0:
		imposto := (150000.0 - salario) * 0.17
		result = imposto
	case salario > 150.000:
		imposto1 := 150000.0 * 0.17
		imposto2 := (salario - 150000.0) * 0.27
		result = imposto1 + imposto2
	}
	return result
}
