package aula2gobases

import "fmt"

/*Um banco deseja conceder empréstimos a seus clientes,
mas nem todos têm acesso a eles.
Ele tem certas regras para determinar quais clientes podem receber empréstimos.
O banco só concede empréstimos a clientes com mais de 22 anos de idade,
que estejam empregados e que estejam em seu emprego há mais de um ano.

Dentro dos empréstimos concedidos, o banco não cobrará juros daqueles que tiverem um salário superior a US$ 100.000.
Você precisa criar um aplicativo que receba essas variáveis e imprima uma mensagem de acordo com cada caso.
📌Dica: seu código deve ser capaz de imprimir pelo menos três mensagens diferentes.*/

func ConcedeEmprestimo(idade int, emprego bool, tempDeContrato float32, salario float32) {
	if idade >= 22 && emprego && tempDeContrato > 1.0 {
		if salario > 100000.0 {
			fmt.Println("Empréstimo concedido sem juros!")
		} else {
			fmt.Println("Empréstimo concedido com juros!")
		}
	} else {
		fmt.Println("Desculpe, pré requisito não cumprido. Não foi possível conceder o empréstimo.")
	}
}
