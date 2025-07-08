package product

import (
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/domain"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/utils"
)

// a interface define os métodos que o serviço de produto deve implementar
type ProductService interface {
	GetAll() (map[int64]domain.Product, error)
	GetByID(id int64) (domain.Product, error)
	GetByParam(priceGt float64) ([]domain.Product, error)
	Create(product domain.RequestProduct) (domain.Product, error)
	Update(id int64, product domain.UpdateProductRequest) (domain.Product, error)
	Patch(id int64, product domain.PatchProductRequest) (domain.Product, error)
	Delete(id int64) error
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
		return nil, utils.ErrNoProductsFound
	}
	return products, nil
}

func (s *productService) GetByID(id int64) (domain.Product, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}
	return product, nil
}

func (s *productService) GetByParam(priceGt float64) ([]domain.Product, error) {
	products, err := s.repository.GetByParam(priceGt)
	if err != nil {
		return nil, utils.ErrNoProductsFound
	}
	return products, nil
}

func (s *productService) Create(product domain.RequestProduct) (domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, utils.ErrNoProductsFound
	}
	newProduct := domain.Product{
		ID:          utils.GenerateUniqueID(products),
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
	createdProduct, err := s.repository.Create(newProduct)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotCreated
	}
	return createdProduct, nil
}

func (s *productService) Update(id int64, product domain.UpdateProductRequest) (domain.Product, error) {
	existing, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	existing.Name = *product.Name
	existing.Quantity = *product.Quantity
	existing.CodeValue = *product.CodeValue
	existing.IsPublished = *product.IsPublished
	existing.Expiration = *product.Expiration
	existing.Price = *product.Price

	updatedProduct, err := s.repository.Update(id, existing)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotUpdated
	}
	return updatedProduct, nil
}

func (s *productService) Patch(id int64, product domain.PatchProductRequest) (domain.Product, error) {
	productToPatch, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}

	switch {
	case product.Name != nil:
		productToPatch.Name = *product.Name
	case product.Quantity != nil:
		productToPatch.Quantity = *product.Quantity
	case product.CodeValue != nil:
		productToPatch.CodeValue = *product.CodeValue
	case product.IsPublished != nil:
		productToPatch.IsPublished = *product.IsPublished
	case product.Expiration != nil:
		productToPatch.Expiration = *product.Expiration
	case product.Price != nil:
		productToPatch.Price = *product.Price
	default:
		return domain.Product{}, utils.ErrNoFieldsToUpdate

	}

	updatedProduct, err := s.repository.Update(id, productToPatch)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotPatched
	}
	return updatedProduct, nil
}

func (s *productService) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return utils.ErrProductNotDeleted
	}
	return nil
}
