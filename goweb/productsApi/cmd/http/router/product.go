package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/cmd/http/handler"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/internal/product"
)

func BuildProductRoutes() http.Handler {
	// Cria um novo roteador Chi
	r := chi.NewRouter()
	// Inicializa o repositório e serviço
	repo, err := product.LoadDatabase()
	if err != nil {
		panic("Failed to load product repository: " + err.Error())
	}
	service := product.NewProductService(repo)
	productHandler := handler.NewProductHandler(service)

	// Registra as rotas de produtos
	r.Group(func(r chi.Router) {
		r.Get("/products", productHandler.GetAll)
		r.Get("/products/{id}", productHandler.GetByID)
		r.Get("/products/search", productHandler.GetByParam)
	})
	r.Group(func(r chi.Router) {
		r.Use(handler.AuthMiddleware)
		r.Post("/products", productHandler.Create)
		r.Put("/products/{id}", productHandler.Update)
		r.Patch("/products/{id}", productHandler.Patch)
		r.Delete("/products/{id}", productHandler.Delete)
	})
	return r
}
