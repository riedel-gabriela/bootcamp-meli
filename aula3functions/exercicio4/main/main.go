package main

/*
Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas para os alunos de um curso.
Para isso, eles precisam gerar uma função que indique o tipo de cálculo que desejam realizar (mínimo, máximo ou médio)
e que retorne outra função e uma mensagem (caso o cálculo não esteja definido) que possa receber um número N de inteiros e retorne o cálculo indicado na função anterior.
Exemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5) retorna 0
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5) retorna 3
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5) retorna 5
*/

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)

}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(operation string) (func(numbers ...float64) float64, error) {
	switch operation {
	case minimum:
		return minValue, nil
	case maximum:
		return maxValue, nil
	case average:
		return averageValue, nil
	default:
		return nil, errors.New("invalid operation")
	}

}

func minValue(numbers ...float64) float64 {
	min := numbers[0]
	for _, number := range numbers {
		min = math.Min(min, number)
	}
	return min
}

func maxValue(numbers ...float64) float64 {
	max := numbers[0]
	for _, number := range numbers {
		max = math.Max(max, number)
	}
	return max
}

func averageValue(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
