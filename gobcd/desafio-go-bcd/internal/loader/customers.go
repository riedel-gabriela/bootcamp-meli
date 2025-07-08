package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// struct do JSON de customer, para trazer os dados do JSON no formato correto para o banco
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

func NewCustomersJSON(path *string) *CustomersJSON {
	return &CustomersJSON{
		path: path,
	}
}

type CustomersJSON struct {
	path *string
}

func (p *CustomersJSON) Load() (customers []internal.Customer, err error) {
	data, err := os.ReadFile(*p.path)
	if err != nil {
		return nil, err
	}
	// criaum slice de CustomerJSON para armazenar os dados do JSON
	var cs []CustomerJSON
	err = json.Unmarshal(data, &cs)
	if err != nil {
		return nil, err
	}
	// converte o slice de CustomerJSON para o slice de internal.Customer
	for _, c := range cs {
		customers = append(customers, internal.Customer{
			Id: c.Id,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: c.FirstName,
				LastName:  c.LastName,
				Condition: c.Condition,
			},
		})
	}

	return customers, nil
}
