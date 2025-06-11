package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func HandleValidationErrors(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
			// Customize error messages
			switch e.Tag() {
			case "required":
				errorMessages[e.Field()] = fmt.Sprintf("'%s' is required.", e.Field())
			case "gte":
				errorMessages[e.Field()] = fmt.Sprintf("'%s' must be non-negative.", e.Field())
			case "ddmmyyyy":
				errorMessages[e.Field()] = fmt.Sprintf("'%s' must be in DD/MM/YYYY format.", e.Field())
			case "min":
				errorMessages[e.Field()] = fmt.Sprintf("'%s' must be at least %s characters long.", e.Field(), e.Param())
			case "max":
				errorMessages[e.Field()] = fmt.Sprintf("'%s' must be at most %s characters long.", e.Field(), e.Param())
			default:
				errorMessages[e.Field()] = fmt.Sprintf("Validation failed for '%s' with tag '%s'.", e.Field(), e.Tag())
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validation failed",
			"errors":  errorMessages,
		})
		return
	}
	// Fallback for other types of errors
	http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
}
