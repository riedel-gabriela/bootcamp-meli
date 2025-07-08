package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

type SaleJSON struct {
	Id        int `json:"id"`
	InvoiceId int `json:"invoice_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func NewSalesJSON(path *string) *SalesJSON {
	return &SalesJSON{
		path: path,
	}
}

type SalesJSON struct {
	path *string
}

func (p *SalesJSON) Load() (sales []internal.Sale, err error) {
	data, err := os.ReadFile(*p.path)
	if err != nil {
		return nil, err
	}
	var ss []SaleJSON
	err = json.Unmarshal(data, &ss)
	if err != nil {
		return nil, err
	}
	for _, s := range ss {
		sales = append(sales, internal.Sale{
			Id: s.Id,
			SaleAttributes: internal.SaleAttributes{
				InvoiceId: s.InvoiceId,
				ProductId: s.ProductId,
				Quantity:  s.Quantity,
			},
		})
	}
	return sales, nil
}
