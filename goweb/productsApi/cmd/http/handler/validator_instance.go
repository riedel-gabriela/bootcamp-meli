package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	customvalidator "github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/utils"
)

// Global validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("ddmmyyyy", customvalidator.ValidateDateFormat)
	log.Println("Validator initialized and custom tags registered.")
}
