package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/domain"
)

type ProductRepository interface {
	GetAll() (map[int64]domain.Product, error)
	GetByID(id int64) (domain.Product, error)
	GetByParam(priceGt float64) ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id int64, product domain.Product) (domain.Product, error)
	Patch(id int64, product domain.Product) (domain.Product, error)
	// Delete(id int64) error
}

type productRepository struct {
	products map[int64]domain.Product
}

func LoadDatabase() (ProductRepository, error) {
	// Lê o arquivo JSON que contém os produtos
	file, err := os.ReadFile("goweb/aula10get/database/products.json")
	// Se houver um erro ao ler o arquivo, retorna um erro formatado
	if err != nil {
		return nil, fmt.Errorf("repository: failed to read file: %w", err)
	}
	// inicia um slice vazio de produtos pois o arquivo json possui um array de produtos
	var products []domain.Product
	if err := json.Unmarshal(file, &products); err != nil {
		return nil, errors.New("repository: failed to unmarshal JSON")
	}
	// transfere para o mapa de produtos, verificar se é necessaria essa etapa mesmo
	productMap := make(map[int64]domain.Product, len(products))
	for _, p := range products {
		productMap[p.ID] = p
	}

	return &productRepository{products: productMap}, nil
}

func (r *productRepository) GetAll() (map[int64]domain.Product, error) {
	if len(r.products) == 0 {
		return nil, fmt.Errorf("repository: no products found")
	}
	return r.products, nil
}

func (r *productRepository) GetByID(id int64) (domain.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return domain.Product{}, fmt.Errorf("repository: product with ID %d not found", id)
	}
	return product, nil
}

func (r productRepository) GetByParam(priceGt float64) ([]domain.Product, error) {
	result := make([]domain.Product, 0)
	for _, product := range r.products {
		if product.Price > priceGt {
			result = append(result, product)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("repository: no products found for price with grather than %f", priceGt)
	}
	return result, nil
}

func (r *productRepository) Create(product domain.Product) (domain.Product, error) {
	r.products[product.ID] = product
	return product, nil
}

func (r *productRepository) Update(id int64, product domain.Product) (domain.Product, error) {
	if _, exists := r.products[id]; !exists {
		return domain.Product{}, fmt.Errorf("repository: product with ID %d not found", id)
	}
	product.ID = id
	r.products[id] = product
	return product, nil
}

func (r *productRepository) Patch(id int64, product domain.Product) (domain.Product, error) {
	original, exists := r.products[id]
	if !exists {
		return domain.Product{}, fmt.Errorf("repository: product with ID %d not found", id)
	}

	// Atualiza apenas os campos enviados (PATCH)
	if product.Name != "" {
		original.Name = product.Name
	}
	if product.Quantity != 0 {
		original.Quantity = product.Quantity
	}
	if product.CodeValue != "" {
		original.CodeValue = product.CodeValue
	}
	if product.Expiration != "" {
		original.Expiration = product.Expiration
	}
	if product.Price != 0 {
		original.Price = product.Price
	}
	original.IsPublished = product.IsPublished // bool zero value = false, então sempre atualiza
	r.products[id] = original
	return original, nil
}
