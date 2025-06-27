package main

import (
	"math"
	"testing"
)

func TestSmallProductPrice(t *testing.T) {
	tests := []struct {
		name     string
		cost     float64
		expected float64
	}{
		{"Custo Positivo", 100.0, 100.0},
		{"Custo Zero", 0.0, 0.0},
		{"Custo Flutuante", 75.50, 75.50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &SmallProduct{cost: test.cost}
			if result := p.Price(); result != test.expected {
				t.Errorf("Para SmallProduct com custo %.2f, esperado %.2f, obtido %.2f", test.cost, test.expected, result)
			}
		})
	}
}

func TestMediumProductPrice(t *testing.T) {
	tests := []struct {
		name     string
		cost     float64
		expected float64 // cost + 3% of cost + 3% of cost
	}{
		{"Custo Positivo", 100.0, 106.0}, // 100 + (100*0.03) + (100*0.03) = 100 + 3 + 3 = 106
		{"Custo Zero", 0.0, 0.0},
		{"Custo Flutuante", 200.0, 212.0},      // 200 + (200*0.03) + (200*0.03) = 200 + 6 + 6 = 212
		{"Custo com Decimais", 150.50, 159.53}, // 150.50 + (150.50*0.03) + (150.50*0.03) = 150.50 + 4.515 + 4.515 = 159.53
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &MediumProduct{cost: test.cost}
			result := p.Price()
			// Usamos uma pequena tolerância para comparações de float64 devido a imprecisões de ponto flutuante
			if math.Abs(result-test.expected) > 0.0001 {
				t.Errorf("Para MediumProduct com custo %.2f, esperado %.2f, obtido %.2f", test.cost, test.expected, result)
			}
		})
	}
}

func TestLargeProductPrice(t *testing.T) {
	tests := []struct {
		name     string
		cost     float64
		expected float64 // cost + 6% of cost + 2500
	}{
		{"Custo Positivo", 100.0, 2606.0},        // 100 + (100*0.06) + 2500 = 100 + 6 + 2500 = 2606
		{"Custo Zero", 0.0, 2500.0},              // 0 + (0*0.06) + 2500 = 2500
		{"Custo Flutuante", 500.0, 3030.0},       // 500 + (500*0.06) + 2500 = 500 + 30 + 2500 = 3030
		{"Custo com Decimais", 250.75, 2765.795}, // 250.75 + (250.75*0.06) + 2500 = 250.75 + 15.045 + 2500 = 2765.795
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &LargeProduct{cost: test.cost}
			result := p.Price()
			if math.Abs(result-test.expected) > 0.0001 {
				t.Errorf("Para LargeProduct com custo %.2f, esperado %.2f, obtido %.2f", test.cost, test.expected, result)
			}
		})
	}
}

func TestProductFactoryInvalidType(t *testing.T) {
	product, err := factory("TipoInvalido", 100.0)

	if err == nil {
		t.Error("Factory para tipo inválido não retornou erro, esperado erro")
	}

	if product != nil {
		t.Error("Factory para tipo inválido retornou um produto, esperado nil")
	}

	expectedErrMsg := "invalid product type"
	if err.Error() != expectedErrMsg {
		t.Errorf("Mensagem de erro incorreta. Esperado '%s', obtido '%s'", expectedErrMsg, err.Error())
	}
}
