package main

import "testing"

func TestDogQtde(t *testing.T) {
	tests := []struct {
		qtde     int
		expected float64
	}{
		{1, 10},
		{2, 20},
		{3, 30},
	}

	for _, test := range tests {
		result := dogQtde(test.qtde)
		if result != test.expected {
			t.Errorf("dogQtde(%d) = %f; want %f", test.qtde, result, test.expected)
		}
	}

}
func TestCatQtde(t *testing.T) {
	tests := []struct {
		qtde     int
		expected float64
	}{
		{1, 5},
		{2, 10},
		{3, 15},
	}

	for _, test := range tests {
		result := catQtde(test.qtde)
		if result != test.expected {
			t.Errorf("catQtde(%d) = %f; want %f", test.qtde, result, test.expected)
		}
	}

}
func TestHamsterQtde(t *testing.T) {
	tests := []struct {
		qtde     int
		expected float64
	}{
		{1, 0.25},
		{2, 0.5},
		{3, 0.75},
	}

	for _, test := range tests {
		result := hamsterQtde(test.qtde)
		if result != test.expected {
			t.Errorf("hamsterQtde(%d) = %f; want %f", test.qtde, result, test.expected)
		}
	}

}
func TestSpiderQtde(t *testing.T) {
	tests := []struct {
		qtde     int
		expected float64
	}{
		{1, 0.15},
		{2, 0.3},
		{4, 0.6},
		{5, 0.75},
	}

	for _, test := range tests {
		result := spiderQtde(test.qtde)
		if result != test.expected {
			t.Errorf("spiderQtde(%d) = %f; want %f", test.qtde, result, test.expected)
		}
	}
}

func TestAnimal(t *testing.T) {
	tests := []struct {
		animal   string
		expected func(int) float64
	}{
		{dog, dogQtde},
		{cat, catQtde},
		{hamster, hamsterQtde},
		{spider, spiderQtde},
	}

	for _, test := range tests {
		result, err := animal(test.animal)
		if err != nil {
			t.Errorf("animal(%s) returned error: %v", test.animal, err)
			continue
		}
		if result == nil {
			t.Errorf("animal(%s) returned nil function", test.animal)
			continue
		}

		if result(1) != test.expected(1) {
			t.Errorf("animal(%s)(1) = %f; want %f", test.animal, result(1), test.expected(1))
		}
	}
}
