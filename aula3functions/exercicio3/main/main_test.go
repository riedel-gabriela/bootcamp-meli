package main

import "testing"

func TestSomaSalario(t *testing.T) {
	tests := []struct {
		minutosTrabalhados int
		categoria          string
		expected           float64
	}{
		{10560, "A", 528001.5},
		{10560, "B", 264001.2},
		{10560, "C", 176000.0},
	}

	for _, test := range tests {
		result := somaSalario(test.minutosTrabalhados, test.categoria)
		if result != test.expected {
			t.Errorf("somaSalario(%d, %s) = %f; want %f", test.minutosTrabalhados, test.categoria, result, test.expected)
		}
	}
}
