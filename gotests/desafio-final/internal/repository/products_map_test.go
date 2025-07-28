package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Aqui o mock pode ser o proprio map

func TestNewProductsMap(t *testing.T) {
	t.Run("NewProductsMap - Success", func(t *testing.T) {
		db := make(map[int]internal.Product)
		repo := NewProductsMap(db)

		assert.NotNil(t, repo)
	})

	t.Run("SearchProducts - Success: no filters", func(t *testing.T) {
		db := map[int]internal.Product{
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
		repo := NewProductsMap(db)

		products, err := repo.SearchProducts(internal.ProductQuery{})
		assert.NoError(t, err)
		assert.Equal(t, db, products)
	})

	t.Run("SearchProducts - Success: with filters", func(t *testing.T) {
		db := map[int]internal.Product{
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
		repo := NewProductsMap(db)

		products, err := repo.SearchProducts(internal.ProductQuery{
			Id: 1,
		})
		assert.NoError(t, err)
		assert.Equal(t, map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       10.0,
				SellerId:    1,
			}},
		}, products)
	})
}
