package aula2gobases

import "fmt"

func PrintContadorLetras(palavra string) {
	fmt.Println("Palavra: ", palavra, "Tamanho: ", len(palavra))
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
}
