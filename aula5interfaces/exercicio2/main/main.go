package main

import (
	"errors"
	"fmt"
)

/*
Algumas lojas de comércio eletrônico precisam criar uma funcionalidade no Go para gerenciar
produtos e retornar o valor do preço total. A empresa tem três tipos de produtos:
Pequeno, Médio e Grande (muitos outros são esperados).

E os custos adicionais são:

Pequeno: apenas o custo do produto
Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500.

O custo de manter o produto em estoque na loja é uma porcentagem do preço do produto.
É necessária uma função factory que receba o tipo de produto e o preço e retorne uma interface Product que tenha o método Price.

Deve ser possível executar o método Price e fazer com que o método retorne o preço total com base no custo do produto
e em quaisquer custos adicionais.
*/

type Product interface {
	Price() float64
}

type SmallProduct struct {
	cost float64
}
type MediumProduct struct {
	cost float64
}
type LargeProduct struct {
	cost float64
}

func (p SmallProduct) Price() float64 {
	return p.cost
}
func (p MediumProduct) Price() float64 {
	return p.cost + (p.cost * 0.03) + (p.cost * 0.03)
}
func (p LargeProduct) Price() float64 {
	return p.cost + (p.cost * 0.06) + 2500
}

type ProductType string

const (
	Small  ProductType = "small"
	Medium ProductType = "medium"
	Large  ProductType = "large"
)

func factory(productType ProductType, cost float64) (Product, error) {
	switch productType {
	case Small:
		return SmallProduct{cost: cost}, nil
	case Medium:
		return MediumProduct{cost: cost}, nil
	case Large:
		return LargeProduct{cost: cost}, nil
	default:
		return nil, errors.New("invalid product type")
	}
}

func main() {
	p1, err := factory(Small, 100.0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Preço do produto Pequeno: %.2f\n", p1.Price()) // Deve ser 100.00
	}

	p2, err := factory(Medium, 100.0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Preço do produto Médio: %.2f\n", p2.Price()) // Deve ser 100 + 3 + 3 = 106.00
	}

	p3, err := factory(Large, 100.0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Preço do produto Grande: %.2f\n", p3.Price()) // Deve ser 100 + 6 + 2500 = 2606.00
	}

	p4, err := factory("Desconhecido", 50.0)
	if err != nil {
		fmt.Println("Erro:", err) // Deve retornar "tipo de produto inválido"
	} else {
		fmt.Printf("Preço do produto Desconhecido: %.2f\n", p4.Price())
	}
}
