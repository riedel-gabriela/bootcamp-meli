package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/cmd/http/router"
)

func main() {
	godotenv.Load()
	r := router.NewRouter()
	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		panic(err)
	}
}
