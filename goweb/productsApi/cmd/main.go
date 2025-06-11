package main

import (
	"net/http"

	"github.com/riedel-gabriela/bootcamp-meli/goweb/productsApi/cmd/http/router"
)

func main() {
	r := router.NewRouter()
	if err := http.ListenAndServe(":3000", r.MapRoutes()); err != nil {
		panic(err)
	}
}
