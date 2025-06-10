package domain

/*
id | number | int | Identificador no conjunto de dados | 15
name | string | string | Nome característico: Cheese - St. Andre
quantity | number | int | Quantidade em estoque: 60
code_value | string | string | Código alfanumérico de característica | S73191A
is_published | boolean | bool | O produto foi publicado ou não: True
expiration | string | string | Data de validade: 12/04/2022
price  | number | float64 | Preço do produto | 50,15
*/

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
