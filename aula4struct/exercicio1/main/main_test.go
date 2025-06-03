package main

import "testing"

func TestGetAll(t *testing.T) {
	tests := []struct {
		expected int
	}{
		{2},
	}

	for _, test := range tests {
		product := Product{}
		product.GetAll()
		if len(products) != test.expected {
			t.Errorf("GetAll() = %d; want %d", len(products), test.expected)
		}
	}

}

func TestSave(t *testing.T) {
	tests := []struct {
		Name        string
		Price       float64
		Description string
		Category    string
		expected    int
	}{
		{"Tablet", 300.00, "Portable tablet with stylus support", "Electronics", 3},
	}

	for _, test := range tests {
		product := Product{
			Name:        test.Name,
			Price:       test.Price,
			Description: test.Description,
			Category:    test.Category,
		}
		product.Save()
		if len(products) != test.expected {
			t.Errorf("Save() = %d; want %d", len(products), test.expected)
		}
	}
}

func TestGetByID(t *testing.T) {
	tests := []struct {
		id       int
		expected string
	}{
		{1, "Laptop"},
		{2, "Smartphone"},
	}

	for _, test := range tests {
		product := Product{}
		p, err := product.GetByID(test.id)
		if err != nil {
			t.Errorf("GetByID(%d) error: %v", test.id, err)
			continue
		}
		if p.Name != test.expected {
			t.Errorf("GetByID(%d) = %s; want %s", test.id, p.Name, test.expected)
		}
	}
}
