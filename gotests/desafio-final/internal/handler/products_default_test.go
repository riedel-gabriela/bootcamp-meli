package handler

import (
	"app/internal"
	"app/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsDefault(t *testing.T) {
	t.Run("Get - Success: no filters", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{}

		expected := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       10.0,
				SellerId:    1,
			}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{
				Description: "Product 2",
				Price:       20.0,
				SellerId:    2,
			}},
		}

		repo.On("SearchProducts", query).Return(expected, nil)

		req := httptest.NewRequest("GET", "/products", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get - Success: with filters", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{Id: 1}

		expected := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       10.0,
				SellerId:    1,
			}},
		}
		repo.On("SearchProducts", query).Return(expected, nil)

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), `"id":1`)
	})

	t.Run("Get - Error: invalid id format", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		// Teste com ID não numérico - não precisa mockar o repositório
		// porque o erro acontece antes de chamar o repositório
		req := httptest.NewRequest("GET", "/products?id=abc", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), `"message":"invalid id"`)
	})

	t.Run("Get - Error: invalid id", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{Id: -1}

		expected := map[int]internal.Product{}

		repo.On("SearchProducts", query).Return(expected, errors.New("invalid id"))

		req := httptest.NewRequest("GET", "/products?id=-1", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), `"message":"invalid id"`)
	})

	t.Run("Get - Error: zero id", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{Id: 0}
		expected := map[int]internal.Product{}

		// Mock para ID zero
		repo.On("SearchProducts", query).Return(expected, errors.New("id cannot be zero"))

		req := httptest.NewRequest("GET", "/products?id=0", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), `"message":"id cannot be zero"`)
	})

	t.Run("Get - Error: internal server error", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{Id: 1}
		expected := map[int]internal.Product{}

		// Mock para simular erro interno (ex: problema de banco de dados)
		repo.On("SearchProducts", query).Return(expected, errors.New("database connection failed"))

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), `"message":"internal server error"`)
	})

	t.Run("Get - Error: unmapped error", func(t *testing.T) {
		repo := new(mocks.ProductsRepositoryMock)
		handler := NewProductsDefault(repo)

		query := internal.ProductQuery{Id: 1}
		expected := map[int]internal.Product{}

		// Mock para simular erro que não está mapeado no switch case (default)
		repo.On("SearchProducts", query).Return(expected, errors.New("some other error"))

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), `"message":"internal server error"`)
	})
}
