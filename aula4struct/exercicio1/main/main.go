package main

import (
	"errors"
	"fmt"
)

func main() {
	product := Product{
		ID:          3,
		Name:        "Tablet",
		Price:       300.00,
		Description: "Portable tablet with stylus support",
		Category:    "Electronics",
	}

	product.Save()

	product.GetAll()

	if p, err := product.GetByID(2); err == nil {
		fmt.Printf("Product found: %+v\n", p)
	} else {
		fmt.Println(err)
	}
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 1500.00, Description: "High performance laptop", Category: "Electronics"},
	{ID: 2, Name: "Smartphone", Price: 800.00, Description: "Latest model smartphone", Category: "Electronics"},
}

func (p *Product) GetAll() {
	fmt.Println("Products:")
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func (p *Product) Save() {
	products = append(products, *p)
	fmt.Printf("Product %s saved successfully!\n", p.Name)
}

func (p *Product) GetByID(id int) (Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found")
}
