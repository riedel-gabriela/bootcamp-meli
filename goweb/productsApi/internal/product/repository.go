package product

import (
	"encoding/json"
	"os"

	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/domain"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/utils"
)

type ProductRepository interface {
	GetAll() (map[int64]domain.Product, error)
	GetByID(id int64) (domain.Product, error)
	GetByParam(priceGt float64) ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id int64, product domain.Product) (domain.Product, error)
	Delete(id int64) error
}

type productRepository struct {
	products map[int64]domain.Product
}

func LoadDatabase() (ProductRepository, error) {
	// Lê o arquivo JSON que contém os produtos
	file, err := os.ReadFile("goweb/productsApi/database/products.json")
	// Se houver um erro ao ler o arquivo, retorna um erro formatado
	if err != nil {
		return nil, utils.ErrCouldntReadFile
	}
	// inicia um slice vazio de produtos pois o arquivo json possui um array de produtos
	var products []domain.Product
	if err := json.Unmarshal(file, &products); err != nil {
		return nil, utils.ErrCouldntUnmarshalJSON
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
		return nil, utils.ErrNoProductsFound
	}
	return r.products, nil
}

func (r *productRepository) GetByID(id int64) (domain.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return domain.Product{}, utils.ErrNoProductsFound
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
		return nil, utils.ErrNoProductsFound
	}
	return result, nil
}

func (r *productRepository) Create(product domain.Product) (domain.Product, error) {
	r.products[product.ID] = product
	return product, nil
}

func (r *productRepository) Update(id int64, product domain.Product) (domain.Product, error) {
	if _, exists := r.products[id]; !exists {
		return domain.Product{}, utils.ErrInvalidProductID
	}
	product.ID = id
	r.products[id] = product
	return product, nil
}

func (r *productRepository) Delete(id int64) error {
	if _, exists := r.products[id]; !exists {
		return utils.ErrInvalidProductID
	}
	delete(r.products, id)
	return nil
}
