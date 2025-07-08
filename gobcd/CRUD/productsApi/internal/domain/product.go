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
	Expiration  string  `json:"expiration"` // Stored as time.Time after parsing
	Price       float64 `json:"price"`
}

type RequestProduct struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int64   `json:"quantity" validate:"required,gte=0"` // Quantity must be required and >= 0
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published" validate:"required"`
	Expiration  string  `json:"expiration" validate:"required,ddmmyyyy"` // Custom validation for DD/MM/YYYY
	Price       float64 `json:"price" validate:"required,gte=0"`         // Price must be required and >= 0
}

type UpdateProductRequest struct {
	ID          int64    `json:"id" validate:"required"`                 // ID is required for update operations (usually from URL path)
	Name        *string  `json:"name" validate:"required,min=2,max=100"` // omitempty: if not present in JSON, don't validate. Min/Max added for example.
	Quantity    *int64   `json:"quantity" validate:"required,gte=0"`
	CodeValue   *string  `json:"code_value" validate:"required"`
	IsPublished *bool    `json:"is_published" validate:"required"`
	Expiration  *string  `json:"expiration" validate:"required,ddmmyyyy"`
	Price       *float64 `json:"price" validate:"required,gte=0"`
}

type PatchProductRequest struct {
	ID          int64    `json:"id" validate:"required"`                            // ID is required for update operations (usually from URL path)
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=2,max=100"` // omitempty: if not present in JSON, don't validate. Min/Max added for example.
	Quantity    *int64   `json:"quantity,omitempty" validate:"omitempty,gte=0"`
	CodeValue   *string  `json:"code_value,omitempty" validate:"omitempty"`
	IsPublished *bool    `json:"is_published,omitempty" validate:"omitempty"`
	Expiration  *string  `json:"expiration,omitempty" validate:"omitempty,ddmmyyyy"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gte=0"`
}
