package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

type InvoiceJSON struct {
	Id         int     `json:"id"`
	CustomerId int     `json:"customer_id"`
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
}

func NewInvoicesJSON(path *string) *InvoicesJSON {
	return &InvoicesJSON{
		path: path,
	}
}

type InvoicesJSON struct {
	path *string
}

func (p *InvoicesJSON) Load() (invoices []internal.Invoice, err error) {
	data, err := os.ReadFile(*p.path)
	if err != nil {
		return nil, err
	}
	// criaum slice de InvoiceJSON para armazenar os dados do JSON
	var ij []InvoiceJSON
	err = json.Unmarshal(data, &ij)
	if err != nil {
		return nil, err
	}
	// converte o slice de InvoiceJSON para o slice de internal.Invoices
	for _, i := range ij {
		invoices = append(invoices, internal.Invoice{
			Id: i.Id,
			InvoiceAttributes: internal.InvoiceAttributes{
				CustomerId: i.CustomerId,
				Datetime:   i.Datetime,
				Total:      i.Total,
			},
		})
	}

	return invoices, nil
}
