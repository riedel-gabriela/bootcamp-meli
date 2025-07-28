package mocks

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type ProductsRepositoryMock struct {
	mock.Mock
}

func (m *ProductsRepositoryMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
