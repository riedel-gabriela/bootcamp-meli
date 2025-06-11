package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/domain"
)

// Valida o formato da data no padrão DD/MM/YYYY
func ValidateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := "02/01/2006" // Go's reference time layout for DD/MM/YYYY

	_, err := time.Parse(layout, dateStr)
	return err == nil
}

// Gera um novo ID único baseado no maior ID existente
func GenerateUniqueID(products map[int64]domain.Product) int64 {
	var maxID int64 = 0
	for id := range products {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
