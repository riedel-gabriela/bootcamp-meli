package utils

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/domain"
)

// Verifica se code_value já existe na lista de produtos
func isCodeValueUnique(products map[int64]domain.Product, codeValue string) bool {
	for _, p := range products {
		if p.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// Valida o formato e valores da data (DD/MM/YYYY)
func isValidDate(date string) bool {
	// verifica se a data está no formato DD/MM/YYYY
	re := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)
	matches := re.FindStringSubmatch(date)
	if len(matches) != 4 {
		return false
	}
	day, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	year, _ := strconv.Atoi(matches[3])
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	// Checagem simples para meses com menos de 31 dias
	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}
	// Fevereiro e ano bissexto
	if month == 2 {
		isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
		if isLeap && day > 29 {
			return false
		}
		if !isLeap && day > 28 {
			return false
		}
	}
	return true
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

// isCreate: true para create, false para update/patch
func ValidateProduct(product *domain.Product, products map[int64]domain.Product, isPatch bool) error {
	// Name
	if !isPatch || (isPatch && strings.TrimSpace(product.Name) != "") {
		if strings.TrimSpace(product.Name) == "" {
			return errors.New("name is required")
		}
	}
	// Quantity
	if !isPatch || (isPatch && product.Quantity != 0) {
		if product.Quantity < 0 {
			return errors.New("quantity cannot be negative")
		}
	}
	// CodeValue
	if !isPatch || (isPatch && strings.TrimSpace(product.CodeValue) != "") {
		if strings.TrimSpace(product.CodeValue) == "" {
			return errors.New("code_value is required")
		}
		if !isCodeValueUnique(products, product.CodeValue) {
			return errors.New("code_value must be unique")
		}
	}
	// Expiration
	if !isPatch || (isPatch && strings.TrimSpace(product.Expiration) != "") {
		if strings.TrimSpace(product.Expiration) == "" {
			return errors.New("expiration date is required")
		}
		if !isValidDate(product.Expiration) {
			return errors.New("expiration date must be in format DD/MM/YYYY and valid")
		}
	}
	// Price
	if !isPatch || (isPatch && product.Price != 0) {
		if math.IsNaN(product.Price) || math.IsInf(product.Price, 0) {
			return errors.New("price must be a valid float64 number")
		}
		if product.Price < 0 {
			return errors.New("price cannot be negative")
		}
	}
	return nil
}
