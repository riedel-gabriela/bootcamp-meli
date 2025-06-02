package main

/*
Uma escola precisa calcular a média (por aluno) de suas notas.
É solicitado que ela gere uma função na qual possam ser passados N números inteiros e que retorne a média.
Não é possível inserir notas negativas.
*/
import "fmt"

func main() {
	fmt.Println("Média: ", media(1, 2, 5, 7, 10, 6, 65))
}

func media(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	soma := 0
	for _, numero := range values {
		soma += numero
	}
	return soma / len(values)
}
