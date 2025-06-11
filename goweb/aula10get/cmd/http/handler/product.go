package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/domain"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/internal/product"
)

type ProductHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetByParam(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	service product.ProductService
}

func NewProductHandler(service product.ProductService) ProductHandler {
	return &productHandler{service: service}
}

func (h *productHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *productHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *productHandler) GetByParam(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	price := query.Get("price")
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Invalid price parameter", http.StatusBadRequest)
		return
	}
	products, err := h.service.GetByParam(priceGt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.RequestProduct
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	err := validate.Struct(newProduct)
	if err != nil {
		HandleValidationErrors(w, err)
		return
	}
	createdProduct, err := h.service.Create(newProduct)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating product: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Product created: %+v\n", createdProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *productHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productToUpdate domain.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productToUpdate); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	productToUpdate.ID = productID
	err = validate.Struct(productToUpdate)
	if err != nil {
		HandleValidationErrors(w, err)
		return
	}

	product, err := h.service.Update(productID, productToUpdate)
	if err != nil {
		if err.Error() == "repository: product with ID "+id+" not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *productHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productToPatch domain.PatchProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productToPatch); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	productToPatch.ID = productID
	err = validate.Struct(productToPatch)
	if err != nil {
		HandleValidationErrors(w, err)
		return
	}

	product, err := h.service.Patch(productID, productToPatch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *productHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(productID)
	if err != nil {
		if err.Error() == "repository: product with ID "+id+" not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
