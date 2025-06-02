package main

import (
	"errors"
	"fmt"
)

/*
Um abrigo de animais precisa calcular a quantidade de alimentos que precisa comprar para seus animais de estimação.
No momento, eles só têm tarântulas, hamsters, cães e gatos, mas esperam poder abrigar muito mais animais. Eles têm os seguintes dados:

Cão 10 kg de comida.
Gato 5 kg de comida.
Hamster 250 g de comida.
Tarântula 150 g de comida.
É solicitado que você:

Implemente uma função Animal que receba como parâmetro um valor de texto com o animal especificado e retorne uma função e uma mensagem (caso o animal não exista).
Uma função para cada animal que calcule a quantidade de comida com base na quantidade do tipo de animal especificado.
Exemplo:

const (
   dog    = "dog"
   cat    = "cat"
)

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/

func main() {
	animalDog, err := animal(dog)
	if err != nil {
		fmt.Println(err)
	}
	animalCat, err := animal(cat)
	if err != nil {
		fmt.Println(err)
	}

	var amount float64
	amount += animalDog(10)
	fmt.Println(amount, "Kg")
	amount += animalCat(10)
	fmt.Println(amount, "Kg")

}

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func animal(animal string) (func(qtde int) float64, error) {
	switch animal {
	case dog:
		return dogQtde, nil
	case cat:
		return catQtde, nil
	case hamster:
		return hamsterQtde, nil
	case spider:
		return spiderQtde, nil
	default:
		return nil, errors.New("non existing animal")
	}
}

func dogQtde(qtde int) float64 {
	return float64(qtde * 10)
}

func catQtde(qtde int) float64 {
	return float64(qtde * 5)
}

func hamsterQtde(qtde int) float64 {
	return float64(qtde) * float64(0.25)
}

func spiderQtde(qtde int) float64 {
	return float64(qtde) * float64(0.15)
}
