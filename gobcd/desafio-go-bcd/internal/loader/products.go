package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewProductsJSON(path *string) *ProductsJSON {
	return &ProductsJSON{
		path: path,
	}
}

type ProductsJSON struct {
	path *string
}

func (p *ProductsJSON) Load() (products []internal.Product, err error) {
	data, err := os.ReadFile(*p.path)
	if err != nil {
		return nil, err
	}
	var ps []ProductJSON
	err = json.Unmarshal(data, &ps)
	if err != nil {
		return nil, err
	}

	// Convert ProductJSON to internal.Product
	for _, p := range ps {
		products = append(products, internal.Product{
			Id: p.Id,
			ProductAttributes: internal.ProductAttributes{
				Description: p.Description,
				Price:       p.Price,
			},
		})
	}
	return products, nil
}
