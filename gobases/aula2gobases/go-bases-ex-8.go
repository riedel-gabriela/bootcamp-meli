package aula2gobases

import (
	"fmt"
)

/*Um funcionário de uma empresa quer saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.
Por outro lado, também é necessário:
Saber quantos de seus funcionários têm mais de 21 anos.
Adicionar um novo funcionário à lista, chamado Federico, que tem 25 anos.
Remover Pedro do mapa.
*/

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func VerificaIdade(nome string) {
	fmt.Println("Idade de ", nome, ": ", employees[nome])
}

func VerificaFaixaEtariaFuncionarios() int {
	contador := 0
	for _, value := range employees {
		if value > 21 {
			contador++
		}
	}
	return contador
}

func AdicionaFuncionario(nome string, idade int) {
	employees[nome] = idade
}

// AdicionaFuncionario("Federico", 25)

func RemoveFuncionario(nome string) {
	delete(employees, nome)
}

// RemoveFuncionario("Pedro")
