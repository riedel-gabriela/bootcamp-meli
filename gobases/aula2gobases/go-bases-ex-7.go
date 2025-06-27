package aula2gobases

import "fmt"

/*
Crie um aplicativo que receba uma variável com o número do mês.

Dependendo do número, imprima o mês correspondente no texto.
Você consegue pensar em maneiras diferentes de resolver isso? Qual delas você escolheria e por quê?
Ex: 7, Julio.

👀Observação: verifique se o número do mês está correto.
*/

func VerificaMesSwitch(numeroDoMes int) {
	switch numeroDoMes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Numero não correspondente. Tente um número entre 1 a 12")
	}
}

func VerificaMesFor(numeroDoMes int) {
	if numeroDoMes < 1 || numeroDoMes > 12 {
		fmt.Println("Numero não correspondente. Tente um número entre 1 a 12")
	}
	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	for i, mesPorExtenso := range meses {
		if i+1 == numeroDoMes {
			fmt.Println(string(mesPorExtenso))
			break
		}
	}
}

func VerificaMesMap(numeroDoMes int) {
	if numeroDoMes < 1 || numeroDoMes > 12 {
		fmt.Println("Numero não correspondente. Tente um número entre 1 a 12")
	}
	meses := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro"}

	fmt.Println(string(meses[numeroDoMes]))
}
