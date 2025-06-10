package product

import (
	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/domain"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/utils"
)

// a interface define os métodos que o serviço de produto deve implementar
type ProductService interface {
	GetAll() (map[int64]domain.Product, error)
	GetByID(id int64) (domain.Product, error)
	GetByParam(priceGt float64) ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id int64, product domain.Product) (domain.Product, error)
	Patch(id int64, product domain.Product) (domain.Product, error)
}

// uma classe productService implementa a interface ProductService
type productService struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (s *productService) GetAll() (map[int64]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) GetByID(id int64) (domain.Product, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *productService) GetByParam(priceGt float64) ([]domain.Product, error) {
	products, err := s.repository.GetByParam(priceGt)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) Create(product domain.Product) (domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	newProduct := domain.Product{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
	if err := utils.ValidateProduct(&newProduct, products, false); err != nil {
		return domain.Product{}, err
	}
	newProduct.ID = utils.GenerateUniqueID(products)
	product, err = s.repository.Create(newProduct)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *productService) Update(id int64, product domain.Product) (domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	if err := utils.ValidateProduct(&product, products, true); err != nil {
		return domain.Product{}, err
	}
	product.ID = id
	updatedProduct, err := s.repository.Update(id, product)
	if err != nil {
		return domain.Product{}, err
	}
	return updatedProduct, nil
}

func (s *productService) Patch(id int64, product domain.Product) (domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	if err := utils.ValidateProduct(&product, products, true); err != nil {
		return domain.Product{}, err
	}
	updatedProduct, err := s.repository.Patch(id, product)
	if err != nil {
		return domain.Product{}, err
	}
	return updatedProduct, nil
}
